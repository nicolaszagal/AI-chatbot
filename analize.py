import sys
import json

import ollama
import openai

transcription_path = sys.argv[1]
question = sys.argv[2]

with open(transcription_path, "r", encoding="utf-8") as f:
    text = f.read()

prompt = f"""
Estás actuando como asistente académico. A continuación tienes la transcripción de una clase universitaria.
Transcripción:
\"\"\"
{text}
\"\"\"

Ahora responde con claridad y precisión a la siguiente pregunta del estudiante:
\"{question}\"
"""

response = ollama.chat(
    model="mistral",
    messages=[
        {"role": "user", "content": prompt}
    ],
)

output = response['message']['content']
print(json.dumps({"answer": output}, ensure_ascii=False))
