## Getting Started on the Receipt Processor

### Prerequisites

Before you begin, ensure you have Go installed on your machine. If Go is not installed, you can download and install it from the https://go.dev/doc/install.

### Setting Up the Project

Follow these steps to set up the project on your local machine:

1. **Initialize the Module**
   Navigate to the root directory of this repository, and run the following command to initialize the module:

   ```bash
   go mod init fetchrewards-go
   ```

   This command will create a `go.mod` file that declares the root of the current module, its module path, and its dependency requirements.

2. **Install Dependencies**
   Install and tidy up the dependencies required for the project by running:

   ```bash
   go mod tidy
   ```

   to ensure that your project has all the necessary dependencies installed and removes any unused ones. Note: it also updates the `go.mod` and `go.sum` files accordingly.

### Building and Running the Application

After setting up the project and its dependencies, you can build and run the application:

1. **Build the Application**
   Compile the project into an executable by running:

   ```bash
   go build
   ```

   This command compiles the packages along with their dependencies and creates an executable file in the current directory.

2. **Run the Application**
   Start the application using:

   ```bash
   go run main.go
   ```

   This command will compile and run the `main.go` file, starting your receipt processor. Ensure that all necessary configuration settings are properly set before running the application.

### Testing the APIs with Jupyter Notebook
The test_demo.ipynb Jupyter Notebook is provided under the root directory of this repository. It includes various tests to ensure the APIs are functioning correctly and returning expected values.

### IDEs Recommended
For developing and running the RESTful API service, I recommend using GoLand, which offers comprehensive support for Go language features and debugging. For running the Jupyter Notebook, Anaconda is my preferred choice as it simplifies package management and deployment of Jupyter environments.
You could also try to use cURL commands or Postman to test the API service.

