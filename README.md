# GO
aprendiendo sobre el lenguaje de programación go para futuros proyectos backend


# apuntes Fundamentos
- es compilado (se compila a código máquina) es tipado (definimos datos) y 
es concurrente (permite ejecutar varias tareas al mismo tiempo)

- Go usa una sintaxis similar a C, pero con algunas diferencias importantes. 



# compilado vs interpretado 

- Compilado (Go, C, C++, Rust): el código se traduce a código máquina una vez, generando un ejecutable (.exe). Después ese ejecutable corre directo. Detectar errores de tipos/sintaxis antes de ejecutar, y corre más rápido.
- Interpretado (Python, JS, PHP): un intérprete lee y ejecuta el código línea por línea en el momento. No genera ejecutable, arranca más rápido pero corre más lento en general.
- ¿Cuál es más fácil de aprender? Interpretado, casi siempre. Podés probar una línea suelta al instante, no hay paso de compilación, y los errores son más claros. Compilado agrega fricción (compilar antes de correr), aunque hoy Go la minimiza con go run que hace todo en un paso. Igual, la lógica que aprendés (variables, loops, funciones) es idéntica en ambos; lo difícil es el lenguaje, no el modelo.



estos fundamentos me serviran como base 
y su concurrencia: permite ejecutar miles de tareas de manera simultánea, consumir muy poca memoria y simplificar la comunicación segura entre procesos sin depender de hilos pesados del sistema operativo
lo cual ayudara en proyectos futuros
