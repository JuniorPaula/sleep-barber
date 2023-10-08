# CS - Dilema do Barbeiro Adormecido

Esta é uma demonstração simples de como resolver o `dilema do Barbeiro Adormecido`, um problema clássico da ciência da computação, que ilustra as complexidades que surgem quando há vários processos do sistema operacional. 

Aqui, temos um número finito de barbeiros, um número finito de assentos em uma sala de espera, um período fixo de tempo que a barbearia fica aberta e clientes vão chegando em intervalos (`aproximadamente`) regulares. Quando um barbeiro não tem nada para fazer, ele verifica a sala de espera a para ver se há novos clientes e, caso haja um ou mais, é feito o corte de cabelo. Caso contrário, o barbeiro vai para a sala de descanço durmir até que um novo cliente chegue. 

Então as regras são as seguintes:
- se não houver clientes, o barbeiro adormece na cadeira.
- um cliente deve acordar o barbeiro se ele estiver dormindo.
- se um cliente chegar enquanto o barbeiro estiver trabalhando, o cliente verificar se há cadeiras disponivéis na sala de espera, se estiver e senta e espera, se não ele vai embora.
- quando o barbeiro termina o corte de cabelo, ele inspeciona a sala de espera para ver se há clientes esperando
e adormece se não houver nenhum.
- a loja pode deixar de aceitar novos clientes na hora de fechar, mas os barbeiros não podem sair até a sala de espera esteja fechada e vazia
- depois que a loja fecha e não há mais clientes na sala de espera, o barbeiro vai para casa.

O **Sleeping Barber** foi originalmente proposto em 1965 pelo pioneiro da ciência da computação Edsger Dijkstra.

O objectivo deste problema, e da sua solução, era deixar claro que, em muitos casos, a utilização de
semáforos (mutexes) não são necessários.

## Referencia
[wikipedia](https://en.wikipedia.org/wiki/Sleeping_barber_problem)

## Linguagem
Para essa demostração foi utilizada a linguagem de **GOLang** pela sua eficiencia em trabalhar com programas concorrentes.