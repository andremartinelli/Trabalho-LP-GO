package main

import (
  "fmt"
)

func result(a *Lista, b *Lista, c *Lista, s Especificacao){
  var total int
  var cont int
  cont = 1
  for aux:=a;aux != nil; aux = aux.proximo{
    if cont == s.vi{
      total = aux.imovel.identificador
    }
    cont++
  }
  cont = 1
  for aux:=b;aux != nil; aux = aux.proximo{
    if cont == s.vj{
      total = total + aux.imovel.identificador
    }
    cont++
  }
  cont = 1
  for aux:=c;aux != nil; aux = aux.proximo{
    if cont == s.vk{
      total =  total + aux.imovel.identificador
    }
    cont++
  }

  fmt.Println(total)
}

func saida(a *Lista, b *Lista, c *Lista, s Especificacao){
  for aux:=a;aux != nil; aux = aux.proximo{
    if aux.proximo == nil{
        fmt.Printf("%d\n",aux.imovel.identificador)
        break
    }
    fmt.Printf("%d, ",aux.imovel.identificador)
  }
  for aux:=b;aux != nil; aux = aux.proximo{
    if aux.proximo == nil{
        fmt.Printf("%d\n",aux.imovel.identificador)
        break
    }
    fmt.Printf("%d, ",aux.imovel.identificador)
  }
  for aux:=c;aux != nil; aux = aux.proximo{
    if aux.proximo == nil{
        fmt.Printf("%d",aux.imovel.identificador)
        break
    }
    fmt.Printf("%d, ",aux.imovel.identificador)
  }
}
