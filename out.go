package main

import (
  // "fmt"
      "os"
      "strconv"
)

func result(a *Lista, b *Lista, c *Lista, s Especificacao){
  file, err := os.Create("result.txt")
    if err != nil {
        return
    }
    defer file.Close()
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
  file.WriteString(strconv.FormatInt(int64(total), 10))
}

func saida(a *Lista, b *Lista, c *Lista, s Especificacao){
  file, err := os.Create("saida.txt")
    if err != nil {
        return
    }
    defer file.Close()

  for aux:=a;aux != nil; aux = aux.proximo{
    if aux.proximo == nil{
        file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
        file.WriteString("\n")
        break
    }
    file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
    file.WriteString(", ")
  }
  for aux:=b;aux != nil; aux = aux.proximo{
    if aux.proximo == nil{
        file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
        file.WriteString("\n")
        break
    }
    file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
    file.WriteString(", ")

  }
  for aux:=c;aux != nil; aux = aux.proximo{
    if aux.proximo == nil{
        file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
        break
    }
    file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
    file.WriteString(", ")

  }
}
