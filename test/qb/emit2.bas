PRINT "Control flow test start"

FOR i = 1 TO 3
    IF i = 2 THEN
        PRINT "Skipping 2"
    ELSE
        PRINT "i =", i
    END IF
NEXT i

x = 1
WHILE x <= 3
    SELECT CASE x
    CASE 1
        PRINT "One"
    CASE 2
        PRINT "Two"
    CASE ELSE
        PRINT "Other"
    END SELECT
    x = x + 1
WEND

PRINT "Done"
