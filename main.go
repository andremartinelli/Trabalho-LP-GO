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
    getAreaTerreno() float32 //caso seja 0 é um terreno, caso seja 1 é uma residência
    getAreaCasa() float32
    getNumeroQuartos() int
}

type tipoTerreno interface{
    areaTerreno() float32
}

type tipoResidencia interface{
    infoResidencia() float32
    infoResidenciaCasa() float32
    areaCasa() float32
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

func (t *Terreno) defineTipoImovel() float32 {
  return (float32)(t.preco)*t.terreno.areaTerreno()*t.solo
}

func (t *Trapezoidal)areaTerreno() float32{
    return t.altura*((t.base1+t.base2)/2.0)
}

func (t *Triangular)areaTerreno() float32{
    return (t.base*t.altura)/2.0
}

func (r *Retangular)areaTerreno() float32{
    return r.lado1*r.lado2
}

func (c *Casa)areaCasa() float32{
    return c.apavimento*(float32)(c.numpavimentos)
}

func (a *Apartamento)areaCasa() float32{
    return 0 //retorna 0 pois não é casa
}

func (r *Residencia) defineTipoImovel() float32 {
  return r.preco*r.residencia.infoResidencia() + r.residencia.infoResidenciaCasa()
}

func (c *Casa)infoResidencia() float32{
  return c.apavimento*(float32)(c.numpavimentos)
}

func (a *Apartamento)infoResidencia() float32{
  return a.aconstruida*(0.9 + (float32)(a.andar)/(float32)(a.totalandares))*a.alazer
}

func (a *Apartamento)infoResidenciaCasa() float32{
  return 0.0
}

func (c *Casa)infoResidenciaCasa() float32{
  return (float32)(c.palivre)*c.alivre
}

func (t *Terreno) getAreaTerreno() float32 {
    return t.terreno.areaTerreno()
  }

func (r *Residencia) getAreaTerreno() float32 {
  return 0 //retorna 0 pois não é terreno
}

func (t *Terreno) getAreaCasa() float32 {
    return 0 //retorna - pois não é casa
  }

func (r *Residencia) getAreaCasa() float32 {
  return r.residencia.areaCasa()
}

func (t *Terreno) getNumeroQuartos() int {
    return 0 //retorna - pois não é residencia
  }

func (r *Residencia) getNumeroQuartos() int {
  return r.quartos
}


func main() {
  fmt.Println("Programa feito por: Andre Martinelli")
    var l *Lista
    var copia *Lista
    var a *Lista
    var b *Lista
    var c *Lista
    var spec Especificacao
    // leitura do arquivo catalogo.txt
    l = lecatalogo(l)
    //leitura do arquivo de operaçoes a serem realizadas
    l = leoperacoes(l)
    //le as especificacoes
    spec = leespec(spec)
    copia = copiaLista(l)
    //cria lista ordenada com os imoveis mais caros
    a = criaListaImoveisCaros(l,a)
    l = copiaLista(copia)
    //Lista dos imóveis mais caros em ordem crescente de preço, respeitando a quantidade especificada
    a = pegaPorcListaCaros(a, spec)
    //lista dos terrenos com menor área
    b = listaMenoresTerrenosArgilosos(l, b)
    //lista dos terrenos com menor área, respeitando a quantidade especificada
    b = pegaPorcListaMenores(b, spec)
    l = copiaLista(copia)
    c = listaAreaCasa(l,c,spec)
    imprimeLista(a)
    result(a,b,c,spec)
    saida(a,b,c,spec)
    }
