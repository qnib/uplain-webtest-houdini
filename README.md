# uplain-webtest-houdini
Webtest server using GPUs

The endpoint `/gpus` spawns `nvidia-smi -L`. Thus, it needs the nvidia driver to be installed and a runtime that supports that.

```
$ curl 127.0.0.1:9999/gpus
webtest_daemon.1.jjncl2qxa9mkymjgasfd9b65o: NVIDIA_VISIBLE_DEVICES=all,HOUDINI_GPU_ENABLED=true
GPU 0: Tesla M60 (UUID: GPU-b85326fc-f486-6b57-87c3-86f035201a5d)
GPU 1: Tesla M60 (UUID: GPU-53f8f4ce-1c4e-035a-98a8-ca7a769d5489)
$ curl 127.0.0.1:9999/gpus
webtest_daemon.2.kecmoyngw9vu2qm98jrbscb35: NVIDIA_VISIBLE_DEVICES=all,HOUDINI_GPU_ENABLED=true
GPU 0: GRID K520 (UUID: GPU-7802b7ad-aabe-2cf5-a871-90d966b79122)
```
