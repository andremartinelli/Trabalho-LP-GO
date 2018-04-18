package main

import (
    "fmt"
    "sort"
)

type Lista struct{
    imovel Imovel
    proximo *Lista
}


func inicializaVazia() *Lista{
    var l Lista
    return &l
}

func insereLista(l *Lista,i Imovel) *Lista{
    var novo Lista
    novo.imovel = i
    novo.proximo = l
    return &novo
}

func insereListaUltimo(l *Lista, i Imovel) *Lista{
  var novo *Lista
  var aux *Lista
  novo.imovel = i
  if l == nil {
    return novo
  }
  for aux = l; aux.proximo != nil; aux = aux.proximo{
    //do nothing
  }
  aux.proximo = novo
  return l
}

func removeLista(l *Lista, ident int) *Lista{
  var aux *Lista
  var aux2 *Lista
  aux2 = l
  for aux = l; aux != nil; aux = aux.proximo{
    if aux.imovel.identificador == ident {
      //remove
      if aux2 == nil { //caso seja o primeiro elemento da lista
        return l.proximo;
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

func imprimeLista(l *Lista){
    var aux *Lista
    var i int
    for aux = l; aux.proximo != nil; aux = aux.proximo{
        fmt.Println(aux.imovel.nome)
        i++
    }
    fmt.Println(i)
}

func criaListaImoveisCaros(l *Lista, a *Lista, s Especificacao) *Lista{
  var tam int
  for aux := l; aux != nil; aux = aux.proximo{
    tam++
  }
  tam--
  aux := l
  ordena := []float64 {}
  for i := 0; i != tam; i++{
    ordena = append(ordena,float64(aux.imovel.imovel.defineTipoImovel()))
    aux = aux.proximo
  }
  sort.Float64s(ordena)
  // porcimv = s.perimocaro*tam/100
  for i:= tam-1; i != -1; i-- {
    for aux := l; aux.proximo != nil; aux = aux.proximo{
      if float64(aux.imovel.imovel.defineTipoImovel()) == ordena[i] {
        a = insereLista(a, aux.imovel)
      }
    }
  }
  // var h int
  // for i := 0; i < tam; i++{
  //   h++
  //   fmt.Printf("%f\n", ordena[i])
  // }
  // fmt.Println(h)
  return a
}
