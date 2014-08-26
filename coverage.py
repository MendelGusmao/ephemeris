#!/usr/bin/env python

import fnmatch
import getopt
import os
import re
import subprocess
import sys

project = "ephemeris"

def initialChecks():
  if "GOPATH" not in os.environ:
    print("Need to set GOPATH")
    sys.exit(1)

def findPath():
  goPath = os.environ["GOPATH"]
  goPathParts = goPath.split(":")
  for goPathPart in goPathParts:
    projectPath = os.path.join(goPathPart, "src", project)
    if os.path.exists(projectPath):
      return projectPath

  return ""

def changePath():
  projectPath = findPath()
  if len(projectPath) == 0:
    print("Project not found")
    sys.exit(1)

  os.chdir(projectPath)

def runCoverReport(allTests):
  print("\n[[ UNIT TESTS ]]\n")

  goPackages = []
  for root, dirnames, filenames in os.walk("."):
    for filename in fnmatch.filter(filenames, "*_test.go"):
      goPackage = project + root[1:]
      goPackages.append(goPackage)

  goPackages = set(goPackages)
  goPackages = list(goPackages)
  goPackages.sort()

  success = True

  for goPackage in goPackages:
    packageName = goPackage.replace("/", "-")

    try:
      subprocess.check_call(["go", "install", goPackage])
      output = subprocess.check_output(["go", "test", "-coverprofile=" + packageName + ".cover.out", "-cover", goPackage])
      matches = re.search(r"coverage: (\d+\.\d+)% of statements", output)

      if matches.group(1) == "100.0" and not allTests:
        os.remove(packageName + ".cover.out")

      sys.stdout.write(output)

    except subprocess.CalledProcessError:
      success = False

  # http://lk4d4.darth.io/posts/multicover/
  mergeCommand = "|".join([
    "echo 'mode: set' > cover-profile.out && cat *.cover.out",
    "grep -v mode:",
    "sort -r",
    "awk '{if($1 != last) {print $0;last=$1}}' >> cover-profile.out"
  ])

  try:
    subprocess.check_call(["sh", "-c", mergeCommand])
    subprocess.check_call(["go", "tool", "cover", "-html=cover-profile.out"])
  except subprocess.CalledProcessError:
    success = False

  # Remove the temporary file created for the
  # covering reports
  try:
    os.remove("cover-profile.out")

    for goPackage in goPackages:
      filename = goPackage.replace("/", "-") + ".cover.out"

      if os.path.isfile(filename):
        os.remove(filename)

  except OSError:
    pass

  if not success:
    print("Errors during the unit test execution")
    sys.exit(1)

###################################################################

def usage():
  print("")
  print("Usage: " + sys.argv[0] + " [-h|--help] [-u|--unit] [-b|--bench] [-n|--integration]")
  print("  Where -h or --help is for showing this usage")
  print("        -a or --all is to show reports for all tests (default: ignore the ones with 100% of coverage)")

def main(argv):
  try:
    opts, args = getopt.getopt(argv, "a", ["all"])

  except getopt.GetoptError as err:
    print(str(err))
    usage()
    sys.exit(1)

  allTests = False

  for key, value in opts:
    if key in ("-a", "--all"):
      allTests = True

    elif key in ("-h", "--help"):
      usage()
      sys.exit(0)

  try:
    initialChecks()
    changePath()
    runCoverReport(allTests)
  except KeyboardInterrupt:
    sys.exit(1)

if __name__ == "__main__":
  main(sys.argv[1:])
