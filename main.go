// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

var choices = []string{"rock", "paper", "scissors", "lizard", "Spock"}

// Mapa de vitórias
var winMap = map[string][]string{
	"rock":     {"scissors", "lizard"},
	"paper":    {"rock", "Spock"},
	"scissors": {"paper", "lizard"},
	"lizard":   {"Spock", "paper"},
	"Spock":    {"scissors", "rock"},
}

// Mapeia cada escolha
var choiceMap = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
	"lizard":   4,
	"Spock":    5,
}

// TAD para teste
type Escolha struct {
	Escolha string `json:"escolha"`
}

func main() {
	// roteador Gin
	r := gin.Default()

	// Rota test GET
	r.GET("/test-connection", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "succesfull handshake"})
	})

	// Rota test POST
	r.POST("/test-body", func(c *gin.Context) {

		var escolha Escolha

		// Bind do body para o TAD Escolha
		if err := c.BindJSON(&escolha); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// retorna o valor recebido
		c.JSON(http.StatusOK, gin.H{"escolha": escolha.Escolha})
	})

	// Rota primeira lógica
	r.POST("/jogar", func(c *gin.Context) {
		// get body
		var jogador struct {
			Escolha string `json:"escolha" binding:"required"`
		}

		if err := c.BindJSON(&jogador); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// escolha aleatória do computador
		rand.Seed(time.Now().UnixNano())
		escolhaComputador := choices[rand.Intn(len(choices))]

		// Achar o vencedor
		resultado := determinarVencedor(jogador.Escolha, escolhaComputador)

		// Retorno do resultado
		c.JSON(http.StatusOK, gin.H{
			"escolha_jogador":    jogador.Escolha,
			"escolha_computador": escolhaComputador,
			"resultado":          resultado,
		})
	})

	// Rota segunda lógica
	r.POST("/jogar2", func(c *gin.Context) {
		var jogador struct {
			Escolha string `json:"escolha" binding:"required"`
		}

		if err := c.BindJSON(&jogador); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rand.Seed(time.Now().UnixNano())
		escolhaComputador := choices[rand.Intn(len(choices))]

		resultado := determinarVencedor(jogador.Escolha, escolhaComputador)

		c.JSON(http.StatusOK, gin.H{
			"escolha_jogador":    jogador.Escolha,
			"escolha_computador": escolhaComputador,
			"resultado":          resultado,
		})
	})

	// Rota terceira lógica
	r.POST("/jogar3", func(c *gin.Context) {
		var jogador struct {
			Escolha string `json:"escolha" binding:"required"`
		}

		if err := c.BindJSON(&jogador); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		rand.Seed(time.Now().UnixNano())
		escolhaComputador := choices[rand.Intn(len(choices))]

		// Obter os valores numericos das escolhas
		valorJogador, ok1 := choiceMap[jogador.Escolha]
		valorComputador, ok2 := choiceMap[escolhaComputador]

		// Verificar se as escolhas são válidas
		if !ok1 || !ok2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Escolha inválida"})
			return
		}

		// Calcular a diferença
		diferenca := (valorJogador - valorComputador + 5) % 5

		// Determinar o resultado
		var resultado string
		switch diferenca {
		case 0:
			resultado = "Empate"
		case 1, 3:
			resultado = "Jogador Vence!"
		case 2, 4:
			resultado = "Computador Vence!"
		}

		// Retornar o resultado
		c.JSON(http.StatusOK, gin.H{
			"escolha_jogador":    jogador.Escolha,
			"escolha_computador": escolhaComputador,
			"resultado":          resultado,
		})
	})

	// Iniciar o servidor
	r.Run(":8080")
}

func determinarVencedor(escolhaJogador, escolhaComputador string) string {
	if escolhaJogador == escolhaComputador {
		return "Empate"
	}

	switch escolhaJogador {
	case "rock":
		return vencedor("scissors", "lizard", escolhaComputador)
	case "paper":
		return vencedor("rock", "Spock", escolhaComputador)
	case "scissors":
		return vencedor("paper", "lizard", escolhaComputador)
	case "lizard":
		return vencedor("Spock", "paper", escolhaComputador)
	case "Spock":
		return vencedor("scissors", "rock", escolhaComputador)
	default:
		return "Escolha inválida"
	}
}

func determinarVencedor2(escolhaJogador, escolhaComputador string) string {
	if escolhaJogador == escolhaComputador {
		return "Empate"
	}

	for _, vitoria := range winMap[escolhaJogador] {
		if escolhaComputador == vitoria {
			return "Jogador Vence!"
		}
	}

	return "Computador Vence!"
}

func vencedor(opcao1, opcao2, escolhaComputador string) string {
	if escolhaComputador == opcao1 || escolhaComputador == opcao2 {
		return "Jogador Vence!"
	}
	return "Computador Vence!"
}
