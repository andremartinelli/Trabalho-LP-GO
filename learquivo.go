package main


import (
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
)

func lecatalogo(l *Lista) *Lista{//lê o arquivo catalogo.txt e salva suas informações
  var floataux float64
  file, err := os.Open("catalogo.txt")
      if err != nil {
          log.Fatal(err)
      }
      defer file.Close()
      scanner := bufio.NewScanner(file)
      for scanner.Scan() {
          if scanner.Text() != "" {
              switch strings.ToLower(scanner.Text()) {
              case "apto":
                  var imovelaux Imovel
                  var residenciaaux Residencia
                  var apartamentoaux Apartamento
                  scanner.Scan()
                  imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  imovelaux.nome = scanner.Text()
                  scanner.Scan()
                  residenciaaux.quartos, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  residenciaaux.vagas, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  apartamentoaux.andar, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
                  apartamentoaux.aconstruida = float32(floataux)
                  scanner.Scan()
                  floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
                  residenciaaux.preco = float32(floataux)
                  scanner.Scan()
                  if scanner.Text() == "S" {apartamentoaux.alazer = 1.15}
                  if scanner.Text() == "N" {apartamentoaux.alazer = 1.0}
                  scanner.Scan()
                  apartamentoaux.totalandares, err = strconv.Atoi(scanner.Text())
                  //terminou de ler o apto
                  residenciaaux.residencia = &apartamentoaux
                  imovelaux.imovel = &residenciaaux
                  l = insereLista(l,imovelaux)

              case "casa":
                  var imovelaux Imovel
                  var residenciaaux Residencia
                  var casaaux Casa
                  scanner.Scan()
                  imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  imovelaux.nome = scanner.Text()
                  scanner.Scan()
                  residenciaaux.quartos, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  residenciaaux.vagas, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  casaaux.numpavimentos, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  casaaux.apavimento = float32(floataux)
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  residenciaaux.preco = float32(floataux)
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  casaaux.alivre = float32(floataux)
                  scanner.Scan()
                  casaaux.palivre, err = strconv.Atoi(scanner.Text())
                  //terminou de ler a casa
                  residenciaaux.residencia = &casaaux
                  imovelaux.imovel = &residenciaaux
                  l = insereLista(l,imovelaux)

              case "triang":
                  var imovelaux Imovel
                  var terrenoaux Terreno
                  var triang Triangular
                  scanner.Scan()
                  imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  imovelaux.nome = scanner.Text()
                  scanner.Scan()
                  if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                  if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                  if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                  scanner.Scan()
                  terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  triang.base = float32(floataux)
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  triang.altura = float32(floataux)
                  //terminou de ler o terreno triangular
                  terrenoaux.terreno = &triang
                  imovelaux.imovel = &terrenoaux
                  l = insereLista(l, imovelaux)

              case "retang":
                  var imovelaux Imovel
                  var terrenoaux Terreno
                  var retang Retangular
                  scanner.Scan()
                  imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  imovelaux.nome = scanner.Text()
                  scanner.Scan()
                  if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                  if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                  if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                  scanner.Scan()
                  terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  retang.lado1 = float32(floataux)
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  retang.lado2 = float32(floataux)
                  //terminou de ler o terreno retangular
                  terrenoaux.terreno = &retang
                  imovelaux.imovel = &terrenoaux
                  l = insereLista(l, imovelaux)

              case "trapez":
                  var imovelaux Imovel
                  var terrenoaux Terreno
                  var trapez Trapezoidal
                  scanner.Scan()
                  imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  imovelaux.nome = scanner.Text()
                  scanner.Scan()
                  if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                  if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                  if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                  scanner.Scan()
                  terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  trapez.base1 = float32(floataux)
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  trapez.base2 = float32(floataux)
                  scanner.Scan()
                  floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                  trapez.altura = float32(floataux)
                  //terminou de ler o terreno trapezoidal
                  terrenoaux.terreno = &trapez
                  imovelaux.imovel = &terrenoaux
                  l = insereLista(l, imovelaux)

              }
          }
      }
      return l
}


func leoperacoes(l *Lista) *Lista{//lê o arquivo atual.txt e salva informações
  var floataux float64
          file, err := os.Open("atual.txt")
              if err != nil {
                  log.Fatal(err)
              }
              defer file.Close()

              scanner := bufio.NewScanner(file)
              for scanner.Scan() {
                if scanner.Text() != ""{
                    switch strings.ToLower(scanner.Text()){
                      case "i":
                        scanner.Scan()
                        if scanner.Text() == "apto" {
                          var imovelaux Imovel
                          var residenciaaux Residencia
                          var apartamentoaux Apartamento
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          residenciaaux.quartos, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          residenciaaux.vagas, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          apartamentoaux.andar, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
                          apartamentoaux.aconstruida = float32(floataux)
                          scanner.Scan()
                          floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
                          residenciaaux.preco = float32(floataux)
                          scanner.Scan()
                          if scanner.Text() == "S" {apartamentoaux.alazer = 1.15}
                          if scanner.Text() == "N" {apartamentoaux.alazer = 1.0}
                          scanner.Scan()
                          apartamentoaux.totalandares, err = strconv.Atoi(scanner.Text())
                          //terminou de ler o apto
                          residenciaaux.residencia = &apartamentoaux
                          imovelaux.imovel = &residenciaaux
                          l = insereLista(l,imovelaux)
                      }
                        if scanner.Text() == "casa" {
                          var imovelaux Imovel
                          var residenciaaux Residencia
                          var casaaux Casa
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          residenciaaux.quartos, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          residenciaaux.vagas, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          casaaux.numpavimentos, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          casaaux.apavimento = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          residenciaaux.preco = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          casaaux.alivre = float32(floataux)
                          scanner.Scan()
                          casaaux.palivre, err = strconv.Atoi(scanner.Text())
                          //terminou de ler a casa
                          residenciaaux.residencia = &casaaux
                          imovelaux.imovel = &residenciaaux
                          l = insereLista(l,imovelaux)
                        }
                        if scanner.Text() == "triang" {
                          var imovelaux Imovel
                          var terrenoaux Terreno
                          var triang Triangular
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                          if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                          if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                          scanner.Scan()
                          terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          triang.base = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          triang.altura = float32(floataux)
                          //terminou de ler o terreno triangular
                          terrenoaux.terreno = &triang
                          imovelaux.imovel = &terrenoaux
                          l = insereLista(l, imovelaux)

                        }
                        if scanner.Text() == "retang" {
                          var imovelaux Imovel
                          var terrenoaux Terreno
                          var retang Retangular
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                          if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                          if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                          scanner.Scan()
                          terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          retang.lado1 = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          retang.lado2 = float32(floataux)
                          //terminou de ler o terreno retangular
                          terrenoaux.terreno = &retang
                          imovelaux.imovel = &terrenoaux
                          l = insereLista(l, imovelaux)
                        }
                        if scanner.Text() == "trapez" {
                          var imovelaux Imovel
                          var terrenoaux Terreno
                          var trapez Trapezoidal
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                          if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                          if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                          scanner.Scan()
                          terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          trapez.base1 = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          trapez.base2 = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          trapez.altura = float32(floataux)
                          //terminou de ler o terreno trapezoidal
                          terrenoaux.terreno = &trapez
                          imovelaux.imovel = &terrenoaux
                          l = insereLista(l, imovelaux)


                        }
                      case "e":
                        scanner.Scan()
                        i := 0
                        i, err =strconv.Atoi(scanner.Text())
                        l = removeLista(l, i)
                      case "a":
                        scanner.Scan()
                        if scanner.Text() == "apto" {
                          var imovelaux Imovel
                          var residenciaaux Residencia
                          var apartamentoaux Apartamento
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          residenciaaux.quartos, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          residenciaaux.vagas, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          apartamentoaux.andar, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
                          apartamentoaux.aconstruida = float32(floataux)
                          scanner.Scan()
                          floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
                          residenciaaux.preco = float32(floataux)
                          scanner.Scan()
                          if scanner.Text() == "S" {apartamentoaux.alazer = 1.15}
                          if scanner.Text() == "N" {apartamentoaux.alazer = 1.0}
                          scanner.Scan()
                          apartamentoaux.totalandares, err = strconv.Atoi(scanner.Text())
                          //terminou de ler o apto
                          residenciaaux.residencia = &apartamentoaux
                          imovelaux.imovel = &residenciaaux
                          l = removeLista(l, imovelaux.identificador)
                          l = insereLista(l,imovelaux)
                      }
                        if scanner.Text() == "casa" {
                          var imovelaux Imovel
                          var residenciaaux Residencia
                          var casaaux Casa
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          residenciaaux.quartos, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          residenciaaux.vagas, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          casaaux.numpavimentos, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          casaaux.apavimento = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          residenciaaux.preco = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          casaaux.alivre = float32(floataux)
                          scanner.Scan()
                          casaaux.palivre, err = strconv.Atoi(scanner.Text())
                          //terminou de ler a casa
                          residenciaaux.residencia = &casaaux
                          imovelaux.imovel = &residenciaaux
                          l = removeLista(l, imovelaux.identificador)
                          l = insereLista(l,imovelaux)
                        }
                        if scanner.Text() == "triang" {
                          var imovelaux Imovel
                          var terrenoaux Terreno
                          var triang Triangular
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                          if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                          if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                          scanner.Scan()
                          terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          triang.base = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          triang.altura = float32(floataux)
                          //terminou de ler o terreno triangular
                          terrenoaux.terreno = &triang
                          imovelaux.imovel = &terrenoaux
                          l = removeLista(l, imovelaux.identificador)
                          l = insereLista(l, imovelaux)

                        }
                        if scanner.Text() == "retang" {
                          var imovelaux Imovel
                          var terrenoaux Terreno
                          var retang Retangular
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                          if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                          if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                          scanner.Scan()
                          terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          retang.lado1 = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          retang.lado2 = float32(floataux)
                          //terminou de ler o terreno retangular
                          terrenoaux.terreno = &retang
                          imovelaux.imovel = &terrenoaux
                          l = removeLista(l, imovelaux.identificador)
                          l = insereLista(l, imovelaux)
                        }
                        if scanner.Text() == "trapez" {
                          var imovelaux Imovel
                          var terrenoaux Terreno
                          var trapez Trapezoidal
                          scanner.Scan()
                          imovelaux.identificador, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          imovelaux.nome = scanner.Text()
                          scanner.Scan()
                          if scanner.Text() == "A" {terrenoaux.solo = 0.9}
                          if scanner.Text() == "G" {terrenoaux.solo = 1.3}
                          if scanner.Text() == "R" {terrenoaux.solo = 1.1}
                          scanner.Scan()
                          terrenoaux.preco, err = strconv.Atoi(scanner.Text())
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          trapez.base1 = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          trapez.base2 = float32(floataux)
                          scanner.Scan()
                          floataux, err = strconv.ParseFloat(scanner.Text(), 64)
                          trapez.altura = float32(floataux)
                          //terminou de ler o terreno trapezoidal
                          terrenoaux.terreno = &trapez
                          imovelaux.imovel = &terrenoaux
                          l = removeLista(l, imovelaux.identificador)
                          l = insereLista(l, imovelaux)
                        }

                    }
                  }
              }
              return l
}

func leespec(s Especificacao) Especificacao{//lê o arquivo espec.txt e salva informações
  var floataux float64
  file, err := os.Open("espec.txt")
      if err != nil {
          log.Fatal(err)
      }
      defer file.Close()

      scanner := bufio.NewScanner(file)
      scanner.Scan()
      s.perimocaro, err = strconv.Atoi(scanner.Text())
      scanner.Scan()
      s.permenarg, err = strconv.Atoi(scanner.Text())
      scanner.Scan()
      floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
      s.alimite= float32(floataux)
      scanner.Scan()
      floataux, err = (strconv.ParseFloat(scanner.Text(), 64))
      s.precolimite = float32(floataux)
      scanner.Scan()
      s.vi, err = strconv.Atoi(scanner.Text())
      scanner.Scan()
      s.vj, err = strconv.Atoi(scanner.Text())
      scanner.Scan()
      s.vk, err = strconv.Atoi(scanner.Text())
      return s
}
