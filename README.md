# Understanding go slices

After tinkering around with slices trying to see its resizing strategy i saw that it uses the doubling strategy that doubles the size of the slice once adding elements at a point where `length > capacity`. I decided to implement a resizing algorithm that uses fibonacci instead of doubling the size so as to avoid wastage of unutilized spaces.



## Drawbacks
* If the slice grows so much it becomes very inefficient

## Benefits
* Minimizes allocation space of contigous blocks of memory when the space required is not that much great