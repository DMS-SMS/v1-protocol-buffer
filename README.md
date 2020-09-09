# proto

## gRPC Compile Command for Python

Use this command in proto directory that contain proto files.

```bash
$ python -m grpc_tools.protoc -I. --python_out=../../python/outing --grpc_python_out=../../python/outing ./outing-student.proto
```

