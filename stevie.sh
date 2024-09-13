
#!/bin/bash

run_make() {
    # Find the Go server process using port 3000
    SERVER_PID=$(lsof -ti tcp:3000 | xargs ps -p | grep '[g]o' | awk '{print $1}')
    
    if [ -n "$SERVER_PID" ]; then
        echo "Killing Go server process using port 3000 (PID: $SERVER_PID)"
        kill -TERM $SERVER_PID  # Try a graceful termination
        sleep 3  # Wait a few seconds to allow the process to terminate

        # If the process is still running, force kill it
        if kill -0 $SERVER_PID 2>/dev/null; then
            echo "Process did not terminate, force killing (PID: $SERVER_PID)"
            kill -9 $SERVER_PID
        fi
    else
        echo "No Go server process found on port 3000"
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
