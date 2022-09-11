# Aufgaben

1. **Erstellen Sie einen Fork von diesem Repository!**
2. **Klonen Sie Ihren Fork, nicht das Original-Repository!**
3. **Reichen Sie Ihre Lösungen per Pull Request ein!**

Die folgende Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## Steckbrief II

`ex1/main.go`: Im letzten Block haben Sie einen Steckbrief mithilfe loser
Variablen definiert (`firstName`, `lastName`, `dayOfBirth` usw.). In dieser
Übung sollen Sie diese Daten zu `struct`s gruppieren. Erstellen Sie die
notwendigen Datenstrukturen und ergänzen Sie die Variable `me` mit Ihren
persönlichen Steckbriefinformationen.

Passen Sie anschliessend die Anzahl Ihrer Geschwister an, indem Sie diese
um eins erhöhen.

## Fibonacci-Zahlen

`ex2/main.go`: Die Fibonacci-Reihe beginnt mit der Folge `1, 1`. Jede weitere
Zahl der Reihe ist definiert durch die Summe des Vorgängers und des
Vor-Vorgängers, also `1, 1, 2, 3`, weil `1+1=2` und `1+2=3` ist, usw. Geht man
von einem 0-basierenden Index aus, kann ein Slice, welches die Fibonacci-Folge
enthält, folgendermassen definiert werden:

- Index `0`: Wert `1`
- Index `1`: Wert `1`
- Index `n`: Wert an Index `n-1` plus Wert an Index `n-2`

In der gegebenen Datei wird ein Slice von fünf Elementen definiert. Die ersten
beiden Werte lauten `1`. Erweitern Sie das Fibonacci-Slice `fibs`
folgendermassen:

1. Indem Sie die verbleibenden drei Werte berechnen und direkt ins bestehende
   Slice schreiben.
2. Indem Sie noch vier weitere Werte berechnen und diese dem Slice anhängen.

## Modulbezeichnungen

`ex3/main.go`: Die Module haben jeweils eine Nummer (z.B. `346`) und eine
Bezeichnung (z.B. "Cloud-Lösungen konzipieren und realisieren"). Erstellen Sie
eine `map` namens `modules`, welche alle nötigen Angaben enthält, um die drei
`fmt.Println()`-Aufrufe unten am Programm die richtigen Informationen ausgeben
zu lassen. Welche Datentypen müssen Sie hierzu verwenden?

Anschliessend nehmen Sie folgende Manipulationen an der `map` vor:

1. Ein Element anhand des Schlüssels entfernen.
2. Ein Element hinzufügen.
3. Ein Element durch ein anderes ersetzen.

Sie können sich Module ausdenken oder im
[Modulbaukasten](https://www.modulbaukasten.ch/) nachschauen, wenn Ihnen die
Modulinformationen nicht geläufig sind.

