
#!/bin/bash

run_make() {
    # Find and kill the process using port 3000
    SERVER_PID=$(lsof -ti tcp:3000)
    if [ -n "$SERVER_PID" ]; then
        echo "Killing process using port 3000 (PID: $SERVER_PID)"
        kill -9 $SERVER_PID  # Force kill to ensure the process is stopped
    fi

    # Start the server by running the make command
    echo "Rebuilding and starting the server..."
    make &
    MAKE_PID=$!
}

# Trap CTRL+C to stop the script gracefully
trap "echo Exiting...; kill $MAKE_PID; exit 0" SIGINT

# Initial make command run
run_make

# Start loop to listen for 'r' key
while true; do
    read -r -n 1 -s input  # Wait for single keypress
    if [[ $input == "r" ]]; then
        echo "Reloading..."
        kill $MAKE_PID  # Kill the running make process
        wait $MAKE_PID 2>/dev/null  # Ensure the process has fully stopped
        run_make  # Restart the make command
    fi
done
