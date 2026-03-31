#!/bin/bash

# Define a function to install all dependencies
install_dependencies() {
  echo "Installing dependencies..."
  npm install
}

# Define a function to run the development server
run_dev_server() {
  echo "Starting development server..."
  npm run start
}

# Define a function to run the production server
run_prod_server() {
  echo "Starting production server..."
  npm run start:prod
}

# Define a function to run tests
run_tests() {
  echo "Running tests..."
  npm run test
}

# Define a function to deploy the application
deploy() {
  echo "Deploying application..."
  npm run deploy
}

# Define a function to clean up the project
clean_up() {
  echo "Cleaning up project..."
  npm run clean
}

# Define the main function
main() {
  case $1 in
    "install")
      install_dependencies
      ;;
    "dev")
      run_dev_server
      ;;
    "prod")
      run_prod_server
      ;;
    "test")
      run_tests
      ;;
    "deploy")
      deploy
      ;;
    "clean")
      clean_up
      ;;
    *)
      echo "Invalid command. Please use one of the following commands: install, dev, prod, test, deploy, clean"
      ;;
  esac
}

# Call the main function with the first argument
main $1