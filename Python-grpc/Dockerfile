FROM python:latest
WORKDIR /project
ADD . /project
RUN pip install -r requirements.txt
EXPOSE 81
CMD ["python","app.py"]