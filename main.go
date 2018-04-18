package main

import (
    // "bufio"
    "fmt"
    // "log"
    // "os"
    // "strings"
    // "strconv"
)

type tipoImovel interface{
    defineTipoImovel() float32
}

type tipoTerreno interface{
    areaTerreno() float32
}

type tipoResidencia interface{
    infoResidencia() float32
}


type Imovel struct{
    nome string
    identificador int
    imovel tipoImovel
}

type Terreno struct{
    preco int //preco m^2 terreno
    solo float32
    terreno tipoTerreno
    //O fatorMultiplicativo varia de acordo com o tipo predominante de solo. Se o solo é
//arenoso, o fator multiplicativo é 0,9; se o solo é argiloso, é 1,3; e se o solo é rochoso, é
//1,1.
}

type Retangular struct{
    lado1 float32
    lado2 float32
}

type Triangular struct{
    base float32
    altura float32
}

type Trapezoidal struct{
    base1 float32
    base2 float32
    altura float32
}

type Residencia struct{
    quartos int
    preco float32
    vagas int
    residencia tipoResidencia
}

type Casa struct{
    numpavimentos int
    apavimento float32
    alivre float32
    palivre int
}

type Apartamento struct{
    andar int
    aconstruida float32
    alazer float32 //será 1.15 caso tenha area de lazer e 1 caso não possua
    totalandares int
}

type Especificacao struct{
  perimocaro int
  permenarg int
  alimite float32
  precolimite float32
  vi int
  vj int
  vk int
}

func (t *Trapezoidal)areaTerreno() float32{
    return t.altura*(t.base1+t.base2)/2
}

func (t *Triangular)areaTerreno() float32{
    return t.base*t.altura
}

func (r *Retangular)areaTerreno() float32{
    return r.lado1*r.lado2
}

func (c *Casa)infoResidencia() float32{
    return c.apavimento*(float32)(c.numpavimentos)+(float32)(c.palivre)*c.alivre
}

func (a *Apartamento)infoResidencia() float32{
    return a.aconstruida*(0.9 + (float32)(a.andar)/(float32)(a.totalandares))*a.alazer
}

func (t *Terreno) defineTipoImovel() float32 {
    return (float32)(t.preco)*t.terreno.areaTerreno()*t.solo
  }

func (r *Residencia) defineTipoImovel() float32 {
  return r.preco*r.residencia.infoResidencia()
}




func main() {
  fmt.Println("Programa feito por: Andre Martinelli")
    var l *Lista
    var a *Lista
    var b *Lista
    var spec Especificacao
    // leitura do arquivo catalogo.txt
    l = lecatalogo(l)
    //leitura do arquivo de operaçoes a serem realizadas
    l = leoperacoes(l)
    //le as especificacoes
    spec = leespec(spec)
    //cria lista ordenada com os imoveis mais caros
    a = criaListaImoveisCaros(l,a)
    //Lista dos imóveis mais caros em ordem crescente de preço, na quantidade especificada
    b = pegaPorcLista(a, spec)
    imprimeLista(b)

    }
