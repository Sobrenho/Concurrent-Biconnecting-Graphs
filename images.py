import pandas as pd
import matplotlib.pyplot as plt

# Load data
data = pd.read_csv("timeTable.csv")

# Initialize data structures for combined graphs
efficiency_data = {}
acceleration_data = {}
elapsed_time_data = {}

# Process each graph and prepare data
for graph_index in data['ID_Grafo'].unique():
    subset = data[data['ID_Grafo'] == graph_index]
    
    seq_time = subset["Tempo Sequencial"].values[0]
    times = [
        seq_time,  # Sequential
        subset["Tempo 1 Thread"].values[0], 
        subset["Tempo 2 Threads"].values[0], 
        subset["Tempo 4 Threads"].values[0], 
        subset["Tempo 8 Threads"].values[0]
    ]
    
    # Store elapsed time data
    elapsed_time_data[graph_index] = times
    
    # Calculate acceleration
    acceleration = [t / seq_time for t in times]
    acceleration_data[graph_index] = acceleration
    
    # Calculate efficiency
    threads = [1, 1, 2, 4, 8]  # Avoid dividing by zero for "Sequential"
    efficiency = [acc / t for acc, t in zip(acceleration, threads)]
    efficiency_data[graph_index] = efficiency

# Plot Elapsed Time vs Number of Threads
plt.figure(figsize=(10, 6))
for graph_index, times in elapsed_time_data.items():
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads'],
        times,
        label=f'Graph {graph_index} - Elapsed Time'
    )
plt.xlabel('Number of Threads')
plt.ylabel('Elapsed Time')
plt.title('Elapsed Time vs Number of Threads')
plt.legend()
plt.grid()
plt.savefig("elapsed_time_combined.png")
plt.close()

# Plot Acceleration vs Number of Threads
plt.figure(figsize=(10, 6))
for graph_index, acceleration in acceleration_data.items():
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads'],
        acceleration,
        label=f'Graph {graph_index} - Acceleration'
    )
plt.xlabel('Number of Threads')
plt.ylabel('Acceleration')
plt.title('Acceleration vs Number of Threads')
plt.legend()
plt.grid()
plt.savefig("acceleration_combined.png")
plt.close()

# Plot Efficiency vs Number of Threads
plt.figure(figsize=(10, 6))
for graph_index, efficiency in efficiency_data.items():
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads'],
        efficiency,
        label=f'Graph {graph_index} - Efficiency'
    )
plt.xlabel('Number of Threads')
plt.ylabel('Efficiency')
plt.title('Efficiency vs Number of Threads')
plt.legend()
plt.grid()
plt.savefig("efficiency_combined.png")
plt.close()
