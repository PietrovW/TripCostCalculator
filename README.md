# TripCostCalculator

*Opis projektu:*
"Kalkulator Kosztów Podróży" to aplikacja api/webowa napisana w języku Go, której zadaniem jest obliczanie całkowitych kosztów podróży. Użytkownik podaje dystans podróży, zużycie paliwa, cenę paliwa, liczbę pasażerów oraz inne dodatkowe koszty, takie jak opłaty za autostrady, parkingi czy przystanki. Aplikacja następnie oblicza całkowity koszt podróży i, opcjonalnie, dzieli ten koszt między pasażerów.

### Funkcjonalności:

1. *Obliczanie kosztów podróży*:
   - Na podstawie podanego dystansu, zużycia paliwa na 100 km oraz ceny paliwa, aplikacja oblicza całkowity koszt paliwa.
   - Dodanie dodatkowych opłat (np. za autostrady, parkingi) do całkowitego kosztu podróży.
   
2. *Podział kosztów*:
   - Opcja podzielenia całkowitego kosztu podróży między pasażerów, co umożliwia obliczenie indywidualnych kosztów dla każdego pasażera.

3. *Historia podróży*:
   - Możliwość zapisania i przeglądania historii podróży oraz obliczonych kosztów dla każdej z nich.

### Schemat działania:

#### Architektura aplikacji:

1. *Main package*: Uruchamia serwer HTTP i obsługuje routing.
2. *Handlers*: Obsługuje zapytania HTTP, takie jak obliczanie kosztów podróży i wyświetlanie historii.
3. *Models*: Definicje struktur danych (np. podróże, koszty).
4. *Services*: Logika obliczeniowa (kalkulacje kosztów podróży).
5. *Database*: Przechowywanie historii podróży w bazie danych (np. SQLite, PostgreSQL).

---

### Kroki działania aplikacji:

1. *Wprowadzenie danych przez użytkownika*: Użytkownik wprowadza dane podróży, takie jak dystans, zużycie paliwa, cena paliwa i dodatkowe opłaty.
2. *Obliczenia*: Aplikacja oblicza koszt paliwa oraz sumuje dodatkowe koszty (opłaty za autostrady, parkingi itp.).
3. *Wyświetlenie wyniku*: Wynik jest zwracany w formacie JSON, a także zapisywany w bazie danych (jeśli taka funkcja jest włączona).
4. *Historia podróży*: Użytkownik może przeglądać historię swoich podróży i związanych z nimi kosztów.




