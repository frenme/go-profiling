# Go Profiling Demo

Демонстрационный проект для изучения инструментов профилирования Go.

## Быстрый старт

```bash
make build
make run
```

## Эндпоинты

- `/cpu` - CPU-интенсивные вычисления
- `/memory` - Выделение памяти  
- `/io` - I/O операции
- `/fibonacci` - Рекурсивные вычисления
- `/benchmark` - Настраиваемые бенчмарки
- `/stats` - Системная статистика

## Профилирование

```bash
# CPU профиль
go tool pprof cpu.prof

# Трассировка
go tool trace trace.out

# Веб-интерфейс pprof
go tool pprof -http=:6060 cpu.prof
```