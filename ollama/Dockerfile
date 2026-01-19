FROM ollama/ollama:latest as builder

# сервер автоматически запускается при pull
RUN bash -c "ollama serve & \
    sleep 10 && \
    ollama pull llama3.2"

FROM ollama/ollama:latest
COPY --from=builder /root/.ollama /root/.ollama
EXPOSE 11434
