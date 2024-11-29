import pandas as pd
import matplotlib.pyplot as plt

# Load data
data = pd.read_csv("timeTable.csv")

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
    
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads'], 
        times,
        label=f'Graph {graph_index} - Elapsed Time'
    )
    
    acceleration = [t / seq_time for t in times]
    
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads'], 
        acceleration,
        label=f'Graph {graph_index} - Acceleration', linestyle='--'
    )
    

    threads = [0, 1, 2, 4, 8] 
    efficiency = [acc / (i if i > 0 else 1) for acc, i in zip(acceleration, threads)]
    
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads'], 
        efficiency,
        label=f'Graph {graph_index} - Efficiency', linestyle=':'
    )

plt.xlabel('Number of Threads')
plt.ylabel('Metrics')
plt.title('Elapsed Time, Acceleration, and Efficiency vs Number of Threads')
plt.legend()
plt.grid()
plt.show()
