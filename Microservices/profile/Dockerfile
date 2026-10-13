FROM python:3.10
WORKDIR /app
COPY source_code/ .
RUN pip install flask
EXPOSE 3001
# Build number shown on the home page (set by CI with --build-arg)
ARG APP_VERSION=dev
ENV APP_VERSION=$APP_VERSION

CMD ["python", "app.py"]
