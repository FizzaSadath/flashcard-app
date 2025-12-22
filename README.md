Flashcard Study App

- Create card decks
- spaced repetition algorithm
- progress tracking
- statistics

SM-2 Algorithm (Super Memo 2)

Variables:

1. Repetitions (n): The "Streak." How many times in a row you have successfully recalled the card.
2. Interval (I): The "Wait Time." How many days until you need to see the card again.
3. Ease Factor (EF): The "Multiplier." A number that represents how easy this specific card is.
    Standard start value: 2.5
    Minimum value: 1.3 (It never gets harder than this).

Grading scale (user input):

When you review a card, you rate yourself from 0 to 5:
5: Perfect response.
4: Correct response after a hesitation.
3: Correct response recalled with serious difficulty.
2: Incorrect response (but you were close).
1: Incorrect response (the correct answer feels familiar).
0: Complete blackout.

Logic flow:

1. Handle the Interval (The Wait Time)
    If you Fail (grade < 3):
        You forgot it. We must reset.
        New Interval: 1 day (Review it tomorrow).
        New Repetitions: 0 (Streak broken).
    If you Succeed (grade >= 3):
        Your memory is getting stronger.
        If Repetitions = 0: New Interval = 1 day.
        If Repetitions = 1: New Interval = 6 days.
        If Repetitions > 1: New Interval = Previous Interval × Ease Factor.
2. Update the Ease Factor (The Difficulty)
    If you got a 5, the EF goes UP (next interval will be much wider).
    If you got a 3, the EF goes DOWN (next interval will grow slowly).
    If you got a 0, the EF drops significantly (but never below 1.3).

Pseudo code:

FUNCTION CalculateReview(current_reps, current_interval, current_ef, user_grade):

    IF user_grade < 3 THEN
        new_reps = 0
        new_interval = 1
   
    ELSE
        new_reps = current_reps + 1

        IF new_reps == 1 THEN
            new_interval = 1
        ELSE IF new_reps == 2 THEN
            new_interval = 6
        ELSE
            new_interval = RoundUp(current_interval * current_ef)
        END IF
    END IF

  
    change_in_ef = (0.1 - (5 - user_grade) * (0.08 + (5 - user_grade) * 0.02))
    new_ef = current_ef + change_in_ef

    IF new_ef < 1.3 THEN
        new_ef = 1.3
    END IF

    RETURN new_reps, new_interval, new_ef