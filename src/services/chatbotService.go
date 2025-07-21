package services

import (
	"fmt"
	"os/exec"
)

func GetAnswer(question string) string {
	answer := map[string]string{
		"¿Cómo estás?":        "¡Estoy bien, gracias por preguntar!",
		"¿Cuál es tu nombre?": "Soy un chatbot creado para ayudarte.",
		"¿Qué puedes hacer?":  "Puedo responder preguntas y ayudarte con información.",
	}

	if answer, ok := answer[question]; ok {
		return answer
	}

	return "Lo siento, no tengo una respuesta para eso."
}

func ChatWithTranscript(transcriptionPath, question string) (string, error) {
	cmd := exec.Command("python", "analize.py", transcriptionPath, question)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing analizE.py")
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(output))
		return "", fmt.Errorf("error generating IA response")
	}
	return string(output), nil
}
