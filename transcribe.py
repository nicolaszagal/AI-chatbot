import sys
import subprocess
import tempfile
import whisper
import json
import os

video_path = sys.argv[1]

# Extract base video name
video_name = os.path.splitext(os.path.basename(video_path))[0]

# Extract audio
temp_audio = tempfile.NamedTemporaryFile(suffix=".wav", delete=False)
temp_audio_path = temp_audio.name
temp_audio.close()

command = [
    'ffmpeg', '-y',
    '-i', video_path,
    '-ar', '16000',
    '-ac', '1',
    '-f', 'wav',
    temp_audio_path
]

print("executing ffmpeg...")
result = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)

if result.returncode != 0:
    print("ffmpeg fail:")
    print(result.stderr.decode("utf-8"))
    sys.exit(1)

print("Audio generated successfully:", temp_audio_path)

# Transcribe audio
try:
    model = whisper.load_model("base", device="cpu")
    result = model.transcribe(temp_audio_path, language="es")

    # Save transcription as a file
    txt_path = os.path.join("transcriptions", f"{video_name}.txt")
    with open(txt_path, "w", encoding="utf-8") as f:
        f.write(result["text"])

    # Send JSON
    print(json.dumps({"transcription": result["text"]}, ensure_ascii=False))
except Exception as e:
    print("Error transcribing with whisper:")
    print(str(e))
    sys.exit(1)