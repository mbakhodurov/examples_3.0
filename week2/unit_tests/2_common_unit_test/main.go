package main

import (
	"log/slog"

	"github.com/mbakhodurov/examples2/week_2/unit_tests/2_common_unit_test/credit_score"
)

func main() {
	client := credit_score.Client{
		Gender:        "male",
		Age:           30,
		Profession:    "engineer",
		Experience:    7,
		AverageSalary: 60000,
	}

	creditScore := credit_score.CalculateCreditScore(client)
	slog.Info("кредитный рейтинг рассчитан", "score", creditScore)
}
