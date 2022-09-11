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
