package main

import (
    "fmt"
    "os"
    "strconv"
    // "io/ioutil"
    // "sort"
)

type Lista struct{
    imovel Imovel
    proximo *Lista
}

func insereLista(l *Lista,i Imovel) *Lista{ //insere um elemenento novo na lista
    var novo Lista
    novo.imovel = i
    novo.proximo = l
    return &novo
}

func insereListaUltimo(l *Lista, i Imovel) *Lista{ //insere um elemento na ultima posição da lista
  var novo *Lista
  var aux *Lista
  novo.imovel = i
  if l == nil {
    return novo
  }
  for aux = l; aux != nil; aux = aux.proximo{
    //do nothing
  }
  aux.proximo = novo
  return l
}

func removeLista(l *Lista, ident int) *Lista{ //remove um elemento da lista
  var aux *Lista
  var aux2 *Lista
  for aux = l; aux != nil; aux = aux.proximo{
    if aux.imovel.identificador == ident {
      //remove
      if aux2 == nil { //caso seja o primeiro elemento da lista
        return l.proximo
      }
      if aux.proximo == nil { //caso seja o ultimo elemento da lista
        aux2.proximo = nil
        return l
      }
      aux2.proximo = aux.proximo
      return l
    }
    aux2 = aux
  }
  return l
}

func imprimeLista(l *Lista){ //imprime os elementos de uma lista num arquivo listaimpressa.txt
  file, err := os.Create("listaimpressa.txt")
    if err != nil {
        return
    }
    defer file.Close()
    var aux *Lista
    for aux = l; aux != nil; aux = aux.proximo{
      file.WriteString(aux.imovel.nome)
      file.WriteString(" - ")
      file.WriteString(strconv.FormatFloat(float64(aux.imovel.imovel.defineTipoImovel()), 'f', 6, 64))
      file.WriteString("\n")
    }
}

func criaListaImoveisCaros(l *Lista, a *Lista, s Especificacao) *Lista{ //cria lista organizada com os imoveis mais caros em ordem crescente
  var tam int
  for aux := l; aux != nil; aux = aux.proximo{
    tam++
  }
  var menor float32
  var atual float32
  var aux2 *Lista
  var aux *Lista
  for i:=0; i < tam; i++{
    menor = 0.0
    atual = 0.0
    for aux = l; aux !=nil; aux = aux.proximo{
      atual = aux.imovel.imovel.defineTipoImovel()
      if menor < atual{
        menor = atual
        aux2 = aux
        // fmt.Println(aux.imovel.nome)
      }
    }
    a = insereLista(a, aux2.imovel)
    l = removeLista(l, aux2.imovel.identificador)
  }
  return a
}
