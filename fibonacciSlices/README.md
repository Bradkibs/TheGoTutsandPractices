# Understanding go slices

After tinkering around with slices trying to see its resizing strategy i saw that it uses the doubling strategy that doubles the size of the slice once adding elements at a point where `length > capacity`. I decided to implement a resizing algorithm that uses fibonacci instead of doubling the size so as to avoid wastage of unutilized spaces.



## Drawbacks
* If the slice grows so much it becomes inefficient

## Benefits
* Minimizes allocation space of contigous blocks of memory when the space required is not that much great

## Typical usecase
* When the slice has fewer items say 50

## Example Comparison
Let's compare the number of resizes for both strategies:

### Doubling Strategy:

Starting with capacity 1: 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024
Number of resizes to reach capacity 1024: 10
### Fibonacci Strategy:

Fibonacci sequence: 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597
Number of resizes to reach capacity 1024: 15