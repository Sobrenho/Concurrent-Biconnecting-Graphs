import pandas as pd
import matplotlib.pyplot as plt

# Read the CSV file
data = pd.read_csv("timeTable.csv")


eficiencia_data = {}
aceleracao_data = {}
tempo_execucao_data = {}

for graph_index in data['grafo'].unique():
    subset = data[data['grafo'] == graph_index]  # Filter rows for the current graph index
    seq_time = subset["Sequencial"].values[0]  # Sequencial time for the current graph
    times = [
        seq_time,
        subset["1 Thread"].values[0],
        subset["2 Threads"].values[0],
        subset["4 Threads"].values[0],
        subset["8 Threads"].values[0],
        subset["16 Threads"].values[0],
    ]

    tempo_execucao_data[graph_index] = times

    aceleracao = [seq_time / t if t != 0 else 0 for t in times] 
    aceleracao_data[graph_index] = aceleracao

    threads = [1, 1, 2, 4, 8, 16]
    eficiencia = [acc / t if t != 0 else 0 for acc, t in zip(aceleracao, threads)]
    eficiencia_data[graph_index] = eficiencia

# Plotting Elapsed Time
plt.figure(figsize=(10, 6))
for graph_index, times in tempo_execucao_data.items():
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads', '16 Threads'],
        times,
        label=f'Grafo {graph_index} - Tempo de Execução'
    )
plt.xlabel('Número de Threads')
plt.ylabel('Tempo de Execução (ms)')
plt.title('Tempo de Execução vs Número de Threads')
plt.legend()
plt.grid()
plt.savefig("elapsed_time_combined.png")
plt.close()

# Plotting Acceleration
plt.figure(figsize=(10, 6))
for graph_index, aceleracao in aceleracao_data.items():
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads', '16 Threads'],
        aceleracao,
        label=f'Graph {graph_index} - Aceleração'
    )
plt.xlabel('Número de Threads')
plt.ylabel('Aceleração')
plt.title('Aceleração vs Número de Threads')
plt.legend()
plt.grid()
plt.savefig("acceleration_combined.png")
plt.close()

plt.figure(figsize=(10, 6))
for graph_index, eficiencia in eficiencia_data.items():
    plt.plot(
        ['Sequencial', '1 Thread', '2 Threads', '4 Threads', '8 Threads', '16 Threads'],
        eficiencia,
        label=f'Grafo {graph_index} - Eficiência'
    )
plt.xlabel('Número de Threads')
plt.ylabel('Eficiência')
plt.title('Eficiência vs Número de Threads')
plt.legend()
plt.grid()
plt.savefig("efficiency_combined.png")
plt.close()
