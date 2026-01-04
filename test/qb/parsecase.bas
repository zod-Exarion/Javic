CLS

PRINT "Enter a number"
INPUT n

SELECT CASE n
    CASE 1
        PRINT "One"
    CASE 2
        PRINT "Two"
    CASE 3
        PRINT "Three"
    CASE ELSE
        PRINT "Other number"
END SELECT

i = 1
WHILE i <= n AND NOT i = 5
    PRINT "i = "; i
    i = i + 1
WEND

PRINT "Done"
END
