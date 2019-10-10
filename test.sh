#!/bin/bash

##############################################################################
#	Runs go fmt/vet/test on each package.
#
#	Usage:		./test.sh {help/...}
##############################################################################

checkSource () {
	# Runs go fmt/vet on source files (vet won't run in loop)
	echo ""
	echo "Running go $1..."
	for I in $(ls); do
		if [ -d $I ]; then
			go $1 $I/*.go
		fi
	done
}

helpText () {
	echo ""
	echo "Runs go fmt/vet/test on each package."
	echo "Usage: ./test.sh {fmt/vet/test}"
	echo ""
	echo "fmt		Runs go fmt on all source files."
	echo "vet		Runs go vet on all source files."
	echo "test		Runs go test on all source files."
	echo "help		Prints help text."
}

if [ $# -eq 0 ]; then
	helpText
elif [ $1 = "fmt" ]; then
	checkSource $1
elif [ $1 = "vet" ]; then
	checkSource $1
elif [ $1 = "test" ]; then
	checkSource $1
elif [ $1 = "help" ]; then
	helpText
else
	helpText
fi
echo ""
