/* O que seria os ponteiros?
- Antes de entender os ponteiros, voce primeiro tem que saber o que cada variavel faz.
As variaveis no seu programa é armazenada em um local especifico na memoria do seu
computador, pense nisso como um numero de casa, apartamento ou ate um numero de ligação
- Ta, eu entendi o que é isso, mas onde quer chegar?
otimo, os ponteiros é aquilo que aponta/mostra pra voce o endereço de uma variavel.
Em vez de conter o valor da variavel diretamente, ele contem o "numero da casa" onde
o valor/variavel está armazenado.
- Entao, é so isso que eu devo aprender?
Negativo, também há o desreferenciamento.
- O que seria o desreferenciamento?
pensa comigo, quando voce tem um ponteiro, pode usar o operador de desrefereciamento que
é o asteristico (*) para acessar ou modificar o real valor que o ponteiro aponta. Voce nao iria
está modificando o ponteiro, e sim o valor da variavel que ele aponta, entende? vou mostrar na pratica
*/

package main

import "fmt"

func main() {
	i := 1                    // iniciei uma variavel normalmente,entao a variavel recebeu um endereço na memoria.
	var p *int = nil          // a variavel p é um ponteiro pra um int, ela está atribuida a nil, o que significa que ela nao esta apontando pra algum endereço.
	p = &i                    //o operador & é usado para obter o endereço de memoria da variavel I, entao, o ponteiro p tem o valor do endereço da variavel I
	*p++                      // o asteristico (*) antes do ponteiro (P) está dizendo pra maquina (Eu quero modificar o valor do i, nao o endereço).
	i++                       // aqui nao precisa do &, pois eu quero modificar o valor e nao o endereço, e muito menos *, pois o i ja tem o real valor, ele nao é um ponteiro.
	fmt.Println(i, p, &i, *p) // o i tem o real valor, o P tem o ponteiro com o endereço, o &I tem o valor do P, e o *P tem o valor do I
} // O console vai retornar = 3 0xc00000a0c8 0xc00000a0c8 3
// primeiro vem o 3, que é o valor, depois em sequencia vem dois endereço da memoria que a variavel está alocada, e novamente o valor do i.
// aprenda que os endereços podem mudar a cada vez que voce executar, pois voce ta encerrando o programa e iniciando novamente, entao ele dar um endereço para variavel
