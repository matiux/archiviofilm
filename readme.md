* Quando si verifica un evento, lanciare il corrispondente metodo, magari in maniera asincrona
* Gestire il verificarsi dell'evento Created - Removed. Si tratta di uno spostamento

# Comandi utili
*Il numero di file in una cartella ricorsivamente:  find . -type f | wc -l

Note: Quando si lancia glide update, bisogna riapplicare la patch al file watcher.go:
Linea 373 - Path:     path2 + " -> " + path1,




61da6203-4417-46a8-b184-b7c6fde27753