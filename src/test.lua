local time = require("time")
print("Ahora:", time.now())
print("Hora formateada:", time.format("%A, %d %b %H:%M:%S %z", time.now()))
print("Dormir un segundo")
time.sleep(1000) -- duerme 1 segundo
print("Listo!")
