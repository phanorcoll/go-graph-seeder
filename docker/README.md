# Running Docker Compose and Using Gremlin Console

Welcome to the **Gremlin Console** setup! This guide will walk you through running your Docker Compose environment and using the Gremlin Console to interact with your graph database. Let's get started!

## Running Docker Compose

1. **Ensure Docker is Installed**: Make sure you have Docker and Docker Compose installed on your machine. You can check this by running:
   ```bash
   docker --version
   docker-compose --version
   ```
2. **Start the Services**: Run the following command to start all services defined in your `docker-compose.yml` file:
   `bash
	docker-compose up -d
	`
   `-d` flag will detach the docker-compose command and run it in the background.

3. **Accessing Logs**: If you want to see the logs for any of the services, you can run:
   `bash
	docker-compose logs -f
	`
   This will give you a real-time view of what's happening inside your containers.

## Using the Gremlin Console

Once your Docker containers are up and running, you can connect to the Gremlin console to interact with your graph database.

1.  **Open a New Terminal Window**: Since you'll be running commands in a different terminal, open a new terminal window.
2.  **Run the Gremlin Console**: Navigate to the directory where your Docker setup is located (the same directory as before) and run:
    ```bash
    docker run -it --rm --network=host tinkerpop/gremlin-console
    ```
    Make sure to execute this command within the Docker folder where your setup is located, since later on we will access the **config** folder from within the console.
3.  **Connect to Your Graph Database**: Once inside the Gremlin console, connect to your remote graph database by typing these two commands:
    ```bash
    # use the remote.yaml file to connect to the local instance of gremlin.
    :remote connect tinkerpop.server conf/remote.yaml
    # switches the console to the remote instance
    :remote console
    ```
4.  **Start Querying**: Now you can start querying your graph! For example, you can retrieve all vertices with a specific label:
    ```bash
    g.V().hasLabel("your_label").limit(10)
    ```
5.  **Exiting the Console**: When you're done, simply type `:exit` or press `Ctrl + D` to leave the Gremlin console.

That's it! You're now set up to use Docker Compose with Gremlin Console. Enjoy exploring your graph database! and don't forget to use **Go graph seeder** to generate data.
