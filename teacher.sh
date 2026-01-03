#!/bin/bash

# Step 1: extract interview number from Buckingham_Place
export KEY_INTERVIEW=$(head -n 179 the-final-cl-test/mystery/streets/Buckingham_Place | tail -n 1 | sed 's/.*#//')

# Step 2: print the environment variable
echo $KEY_INTERVIEW

# Step 3: print the contents of the interview
cat "the-final-cl-test/mystery/interviews/interview-$KEY_INTERVIEW"

# Step 4: print the MAIN_SUSPECT env variable
echo $MAIN_SUSPECT
