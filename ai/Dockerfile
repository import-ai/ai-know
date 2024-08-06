FROM python:3.10
RUN python3 -m pip install "torch~=2.3.1" --index-url https://download.pytorch.org/whl/cpu
COPY requirements.txt /requirements.txt
RUN python3 -m pip install -r /requirements.txt \
    -i https://mirrors.ustc.edu.cn/pypi/web/simple \
    -i https://pypi.tuna.tsinghua.edu.cn/simple
