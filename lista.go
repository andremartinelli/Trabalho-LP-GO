package main

import (
    "fmt"
    "os"
    "strconv"
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
  var novo Lista
  var aux *Lista
  novo.imovel = i
  if l == nil {
    return &novo
  }
  for aux = l; aux.proximo != nil; aux = aux.proximo{
  }
    aux.proximo = &novo
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
  fmt.Println("Imprimindo (função imprimeLista)")
  file, err := os.Create("listaimpressa.txt")
    if err != nil {
        return
    }
    defer file.Close()
    var aux *Lista
    for aux = l; aux != nil; aux = aux.proximo{
      file.WriteString(strconv.FormatInt(int64(aux.imovel.identificador), 10))
      file.WriteString(" - ")
      file.WriteString(aux.imovel.nome)
      file.WriteString(" - ")
      file.WriteString(strconv.FormatFloat(float64(aux.imovel.imovel.defineTipoImovel()), 'f', 6, 64))
      file.WriteString("\n")
    }
}

func copiaLista(l *Lista) *Lista{
  var a *Lista
  for aux:=l; aux != nil; aux = aux.proximo{
    a = insereLista(a, aux.imovel)
  }
  return a
}

func criaListaImoveisCaros(l *Lista, a *Lista) *Lista{ //cria lista organizada com os imoveis mais caros em ordem crescente
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
      if atual == menor{//caso seja igual
        if aux.imovel.identificador < aux2.imovel.identificador{ //pega o com o menor identificador e seta como menor
          menor = atual
          aux2 = aux
        }
      }
      if menor < atual{
        menor = atual
        aux2 = aux
      }
    }
    a = insereListaUltimo(a, aux2.imovel)
    l = removeLista(l, aux2.imovel.identificador)
  }
  return a
}

func pegaPorcListaCaros(a *Lista, s Especificacao) *Lista{//retorna lista dos imóveis mais caros em ordem crescente de preço, na quantidade especificada em spec.txt
  var por float32
  var tam int
  var novalista *Lista
  for aux := a; aux != nil; aux = aux.proximo{
    tam++
  }
  por = float32(s.perimocaro)*float32(tam)/100.0
  b := int(por)
  i := 0
  for aux := a; i < b; aux = aux.proximo{
    novalista = insereLista(novalista, aux.imovel)
    i++
  }
  return novalista
}


func listaMenoresTerrenosArgilosos(l *Lista, a *Lista) *Lista{//retorn lista com menores terrenos
  var tam int
  for aux := l; aux != nil; aux = aux.proximo{
    if aux.imovel.imovel.getAreaTerreno() != 0 && aux.imovel.imovel.getTipoTerreno() == 1.3{tam++}
  }
  var menor float32
  var atual float32
  var aux2 *Lista
  var aux *Lista
  for i:=0; i < tam; i++{
    menor = 0.0
    atual = 0.0
    for aux = l; aux !=nil; aux = aux.proximo{
      if aux.imovel.imovel.getAreaTerreno() != 0 && aux.imovel.imovel.getTipoTerreno() == 1.3{
        atual = aux.imovel.imovel.getAreaTerreno()
        if atual == menor{//caso seja igual
          if aux.imovel.identificador < aux2.imovel.identificador{//pega o com o menor identificador e seta como menor
            menor = atual
            aux2 = aux
          }
        }
        if menor < atual{
          menor = atual
          aux2 = aux
        }
      }
    }
    a = insereLista(a, aux2.imovel)
    l = removeLista(l, aux2.imovel.identificador)
  }
  return a
}

func pegaPorcListaMenores(a *Lista, s Especificacao) *Lista{
  var por float32
  var tam int
  var novalista *Lista
  for aux := a; aux != nil; aux = aux.proximo{
    tam++
  }
  por = float32(s.permenarg)*float32(tam)/100.0
  b := int(por)
  i := 0
  for aux := a; i < b; aux = aux.proximo{
    novalista = insereLista(novalista, aux.imovel)
    i++
  }
  return novalista
}


func listaAreaCasa(l *Lista, a *Lista, s Especificacao) *Lista{
  var tam int
  for aux := l; aux != nil; aux = aux.proximo{
    if aux.imovel.imovel.getAreaCasa() > s.alimite && aux.imovel.imovel.defineTipoImovel() < s.precolimite{tam++}
  }
  var menor int
  var atual int
  var aux2 *Lista
  var aux *Lista
  for i:=0; i < tam; i++{
    menor = 0
    atual = 0
    for aux = l; aux !=nil; aux = aux.proximo{
      if aux.imovel.imovel.getAreaCasa() > s.alimite && aux.imovel.imovel.defineTipoImovel() < s.precolimite{
        atual = aux.imovel.imovel.getNumeroQuartos()
        if atual == menor{//caso seja igual
          if aux.imovel.identificador > aux2.imovel.identificador{ //pega o com o menor identificador e seta como menor
            menor = atual
            aux2 = aux
          }
        }
        if menor < atual{
          menor = atual
          aux2 = aux
        }
      }
    }
    a = insereListaUltimo(a, aux2.imovel)
    l = removeLista(l, aux2.imovel.identificador)
  }
  return a
}
