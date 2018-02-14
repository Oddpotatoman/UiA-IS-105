

| Binære tall | Hexadesimaltall | Desimaltall |
| --- | --- | --- |
| 1101 | 0xD | 13 |
| 110111101010 | DEA | 3562 |
| 1010111100110100 | 0XAF34 | 44852 |
| 1111111111111111 | FFFF | 65535 |
| 10001011110001010 | 1178A | 71562 |

Oppgave 1)

**Oppgave A)**

**Hexadesimale tall til binære tall:**

For å gå fra hexadesimale tall til binære tall er det enkleste å se hvert enkelt siffer i det hexadesimale tallet og finne det samme binære tallet. Etter det slår vi sammen de binære tallene i nøyaktig samme rekkefølge de stod i det hexadesimale tallet.  Eksempel F5A3:

| Heksadesimal siffer | Binært tall | Desimalt tall |
| --- | --- | --- |
| F | 1111 | 15 |
| 5 | 0101 | 5 |
| A | 1010 | 10 |
| 3 | 0011 | 3 |

Når vi har funnet ut de binære tallene legger vi de sammen i rekkefølge, som følgende. 1111+0101+1010+0011 = 11110101 1010011. Vi pleier å dele opp i grupper på 8 siffer for å kunne lese tallene bedre.

**Fra binære tall til hexadesimal:**

Dette er nesten samme som fra hexadesimale tall til binære tall. Forskjellen er at vi ikke tar hvert enkelt siffer fra det binære tallet, men vi gjør det om til firesifret  binært tall. Vi begynner alltid bakerst i det binære tallet, helt til høyre.  Om vi ikke har et firesifret  tall til slutt fyller vi på med 0´er avhengig av hvor mange vi trenger for å få firesifret tall til slutt. Eksempel 11000101

| Binært tall | Hexadesimal tall | Desimalt tall |
| --- | --- | --- |
| 0101 | 5 | 5 |
| 1100 | C | 12 |

Det som er viktig nå er at rekkefølgen følger den binære rekkefølgen. Så her får vi C+5= C5.









**Desimaltall til binære tall:**

Når vi skal gå fra desimaltall til binære tall, deler vi tallet på 2, får vi rest skal vi skrive 1, om ikke vi får rest skriver vi 0.  Vi begynner alltid på toppen og går ned til heltallsdivisjonen er 0. Eksempel 53

| **Verdi** | **Utregning** | **Rest** |
| --- | --- | --- |
| **53** | **53/2=1** | **1** |
| **26** | **26/2=0** | **0** |
| **13** | **13/2=1** | **1** |
| **6** | **6/2=0** | **0** |
| **3** | **3/2=1** | **1** |
| **1** | **1 / 2 = 1** | **1** |



Når vi har gjort dette, leser vi tallet nedenfra og opp slik at vi får 110101.

**Binære tall til desimaltall:**

For å få fra binære tall til desimaltall bruker vi formelen Verdi = siffer\*grunntall^posisjon. Eksempel 111100

| Posisjon | Siffer | Utregning |
| --- | --- | --- |
| 0 | 0 | 0x2^0=0 |
| 1 | 0 | 0x2^1=0 |
| 2 | 1 | 1x2^2=4 |
| 3 | 1 | 1x2^3=8 |
| 4 | 1 | 1x2^4=16 |
| 5 | 1 | 1x2^5=32 |

Når vi da har funnet desimaltallene plusser vi de bare sammen 32+16+8+4=60=111100















**Oppgave B)**

**Hexadesimale tall til desimal**

I det hexadesimale tallsystem så er grunntallet 16, 0-9 er i tall og 10-15 er bokstaver. A = 10  B = 11 C = 12 D = 13 E = 14 F = 15. Her bruker vi også formelen verdi=siffer\*grunntall^posisjon. Eksempel DAAF

| Posisjon | Siffer | Utregning |
| --- | --- | --- |
| 0 | F | 15\*16^0=15 |
| 1 | A | 10\*16^1=160 |
| 2 | A | 10\*16^2=2560 |
| 3 | D | 10\*16^3=53248 |

15+160+2560+53248=55983=DAAF

**Desimal til hexaldesimal:**

Utregningen fra desimal til hexadesimal følger samme konvertering som fra desimaltall til binært tall. Forskjellen er at vi deler på 16 og ikke 2. Får vi rest, så skriver vi ned resten. Får vi ikke rest, skriv 0.  Eksempel 520

| Verdi | Utregning | Rest | Hexadesimal tall |
| --- | --- | --- | --- |
| 520 | 520/16=32 32\*16=512 | 8 | 8 |
| 32 | 32/16=2 | 0 | 0 |
| 2 | 2/16 | 2 | 2 |

Det som er viktig å huske på er når vi tar 520/16=32  må vi gange resultatet med 16 igjen. Da får vi 32\*16=512. Differansen mellom tallet vi begynte med 520 og 512 er 8, da får vi en rest på 8. Det skal gjøres helt til vi går i 0.
