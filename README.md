# Strukturen, Slices und Maps

Im [ersten Teil](https://code.frickelbude.ch/m346/go-1-vars-types-output) haben
wir _primitive Datentypen_ betrachtet. Im zweiten Teil geht es um
_zusammengesetzte Datentypen_ (engl. _Compound Data Types_), welche sich aus aus
primitiven und anderen zusammengesetzten Datentypen zusammensetzen können.

## Strukturen

Wir haben gesehen, dass `byte` bloss ein Alias für `uint8` ist, was mithilfe von
`type` folgendermassen bewerkstelligt wird:

```go
type byte uint8
```

Man kann also beliebige bestehende Datentypen nehmen und ihnen eine zusätzliche
Bezeichnung geben, beispielsweise:

```go
type float32 degrees
type int16 year
type rune sign
```

Mehrere Variablen gleichen oder unterschiedlichen Typs können zu sogenannten
_Strukturen_ ([Spec](https://go.dev/ref/spec#Struct_types)) zusammengefügt
werden:

```go
struct {
    [Element] [Datentyp]
    ...
}
```

Mit einer `struct` wird ein neuer Datentyp definiert. Mithilfe des
`type`-Schlüsselworts ([Spec](https://go.dev/ref/spec#Type_declarations)) kann
diesem neuen Typ eine Bezeichnung gegeben werden:

```go
type [Bezeichnung] struct {
    [Element] [Datentyp]
    ...
}
```

Eine Struktur ist _heterogen_, d.h. es können darin Variablen verschiedener
Datentypen gespeichert werden. In diesem Beispiel werden Informationen von einem
Steckbrief in einer Struktur abgelegt:

```go
type Person struct {
    FirstName    string
    LastName     string
    DayOfBirth   byte
    MonthOfBirth byte
    YearOfBirth  int16
}
```

Für die Gross- und Kleinschreibung sollen bis auf Weiteres folgende Regeln
gelten:

- Neue Datentypen wie Strukturen werden mit grossem Anfangsbuchstaben
  geschrieben, z.B.  `Person` oder `Address`.
- Elemente von Strukturen werden ebenfalls mit grossem Anfangsbuchstaben
  geschrieben, z.B. `FirstName` oder `LastName`.
- Konkrete Variablen werden mit kleinem Anfangsbuchstaben geschrieben, also
  `person` oder `address`.

Eine Struktur kann auch aus anderen Strukturen zusammengesetzt werden. So lassen
sich die Informationen eines Steckbriefs gruppieren:

```go
type FullName struct {
    FirstName string
    LastName  string
}

type BirthDate struct {
    DayOfBirth   byte
    MonthOfBirth byte
    YearOfBirth  int16
}

type Person struct {
    Name FullName
    Born BirthDate
}
```

Eine Struktur kann als Variable deklariert werden:

```go
var myName FullName
var myBirthDate BirthDate
```

Die Initialisierung der Struktur erfolgt mit folgender Syntax:

```go
var myName FullName = FullName{
    FirstName: "Patrick",
    LastName:  "Bucher",
}
var myBirthDate BirthDate = BirthDate{
    DayOfBirth:   24,
    MonthOfBirth: 6,
    YearOfBirth:  1987,
}
```

Beachten Sie, dass am Ende jeder Zeile ein Komma stehen muss! (Das hat Vorteile
beim Umsortieren der Angaben, und in der Versionskontrolle sieht man dadurch
beim Hinzufügen einer Angabe keine Änderung auf der vorherigen Zeile durch das
Hinzufügen des Kommas.)

Der Typ bei der Deklaration (links vom `=`) kann weggelassen werden:

```go
var teacher = Person{
    Name: myName,
    Born: myBirthDate,
}
```

Da die Reihenfolge der Elemente festgelegt ist, kann deren Name weggelassen
werden:

```go
var teacher = Person{
    myName,
    myBirthDate,
}
```

Lesender und schreibender Zugriff auf einzelne Elemente der Struktur ist
mithilfe des Punkt-Operators möglich:

```go
fmt.Println("Teacher's last name:", teacher.Name.LastName)

teacher.Name.FirstName = "Padraigh"
fmt.Println("Teacher's Irish name:", teacher.Name.FirstName)
```

Ausgabe:

    Teacher's last name: Bucher
    Teacher's Irish name: Padraigh

### Einbetten von Strukturen

Strukturen können aus anderen Strukturen bestehen. Beim Zugriff auf die
Unterelemente muss der Zugriff entsprechend über mehrere Schritte erfolgen, also
z.B. `teacher.Name.LastName` im vorherigen Beispiel.

Verzichtet man beim Einbetten von Unterstrukturen auf einen Namen, können die
Unterelemente direkt angesprochen werden:

```go
type Teacher struct {
    FullName
    BirthDate
    TeachesModule string
}

pbucher := Teacher{
    myName,
    myBirthDate,
    "Modul 346",
}

fmt.Println("Teacher's full name:", pbucher.FirstName, pbucher.LastName)
fmt.Printf("Teacher's birth date: %d.%d.%d\n",
    pbucher.DayOfBirth, pbucher.MonthOfBirth, pbucher.YearOfBirth)
```

Ausgabe:

    Teacher's full name: Patrick Bucher
    Teacher's birth date: 24.6.1987

### Ausgabe von Strukturen

Mit der Formatangabe `%v` kann eine Struktur direkt mit allen ihren Elementen
ausgegeben werden. Mit der Formatangabe `%q` erfolgt die Ausgabe in Go-Syntax:

```go
fmt.Printf("%v\n", pbucher)
fmt.Printf("%q\n", pbucher)
```

Ausgabe:

    {{Patrick Bucher} {24 6 1987} Modul 346}
    {{"Patrick" "Bucher"} {'\x18' '\x06' '߃'} "Modul 346"}

(Offenbar wurden die Angaben des Geburtsdatums auf der zweiten Zeile als
hexadezimale bzw. Unicode-Zeichen interpretiert. Die generische Ausgabe mit `%q`
ist zwar sehr einfach, aber nicht in jedem Fall sehr hilreich.)

## Slices

_Slices_ ([Spec](https://go.dev/ref/spec#Slice_types)) sind _homogen_, d.h. es
können darin Werte des gleichen Datentyps abgespeichert werden. Im Gegensatz zu
einer Struktur können das (theoretisch) beliebig viele Werte sein; die
verfügbare Menge an Arbeitsspeicher ist die einzige Grenze.

Strenggenommen stellen Slices nur eine Sicht auf _Arrays_
([Spec](https://go.dev/ref/spec#Array_types)) dar, welche die eigentlichen Daten
beinhalten. Für unsere Zwecke ist aber diese Entscheidung nicht notwendig. Darum
soll hier nur von Slices die Rede sein.

Ein Slice wird wie eine Variable deklariert. Dem Typ geht aber ein eckiges
Klammernpaar voraus:

```go
var name string    // a single string
var names []string // a slice of strings
```

Bei der Initialisierung kann wiederum auf die Typangabe links vom `=` verzichtet
werden, da diese rechts davon angegeben wird:

```go
var days = []string{"Mo", "tu", "We", "Th", "Fr", "Sa"}
```

Beim `days`-Slice ging offenbar der Sonntag vergessen, der mithilfe von `append`
ans Ende des Slices angefügt werden kann:

```go
days = append(days, "Su")
```

Achtung: `append` gibt eine Referenz auf ein neues Slice zurück, weswegen der
Rückgabewert wiederum abgespeichert werden muss.

Das zweite Element `"tu"` wurde im Gegensatz zu den anderen Elementen
kleingeschrieben. Dies soll angepasst werden, indem das Element mithilfe des
0-basierenden Index überschrieben wird:

```go
days[1] = "Tu"
```

Das Slice sollte nun die abgekürzten Wochentage enthalten. Deren Anzahl kann mit
der eingebauten `len()`-Funktion ermittelt werden:

```go
fmt.Println(days)
fmt.Println(len(days))
```

Ausgabe:

    [Mo Tu We Th Fr Sa Su]
    7

Das erste Element ist an Index `0` zu finden. Für den Zugriff auf das letzte
Element kann die Länge vom Slice abzüglich eins verwendet werden:

```go
firstDay := days[0]
lastDay := days[len(days)-1]
fmt.Println("from", firstDay, "to", lastDay)
```

Ausgabe:

    from Mo to Su

### Slice-Ausschnitte und Slicing-Syntax

Mit der namensgebenden _Slicing-Syntax_ können Ausschnitte aus dem Slice
ermittelt werden. Hierzu wird die Untergrenze _inklusiv_, die Obergrenze
_exklusiv_ angegeben, sodass die Obergrenze von der Untergrenze subtrahiert die
gleich der Länge des neuen Slices ist:

```go
workdays := days[0:5]
weekend := days[5:7]
fmt.Println(workdays, len(workdays))
fmt.Println(weekend, len(weekend))
```

Ausgabe:

    [Mo Tu We Th Fr] 5
    [Sa Su] 2

### Exkurs: Länge und Kapazität (optional)

Mithilfe der eingebauten `make`-Funktion lassen sich verschiedene
Datenstrukturen erzeugen, z.B. Slices. Hierzu wird ein Typ und die initiale
Grösse des Slices angegeben:

```go
var numbers = make([]int, 0)
var moreNumbers = make([]int, 3)
```

Ein Slice hat neben einer Länge (`len()`) auch eine _Kapazität_ (`cap()`) für
weitere Elemente:

```go
fmt.Println(numbers, len(numbers), cap(numbers))
fmt.Println(moreNumbers, len(moreNumbers), cap(moreNumbers))
```

Ausgabe:

    [] 0 0
    [0 0 0] 3 3

Das erste Slice (`numbers`) enthält keine Werte und hat darum die Länge und
Kapazität 0. Das zweite Slice besteht aus drei Werten und hat darum die Länge
und Kapazität 3.

Der Unterschied zwischen Länge und Kapazität wird erst dann klar, wenn man
mithilfe der Slicing-Syntax auf einen Unterbereich zugreift:

```go
var extract = moreNumbers[0:2]
fmt.Println(extract, len(extract), cap(extract))
```

Ausgabe:

    [0 0] 2 3

Die Länge beträgt jetzt nur noch `2`, doch die Kapazität beträgt nach wie vor
`3`, weil das zugrundeliegende Slice (bzw. Array) noch einen weiteren Wert
aufnehmen kann.
