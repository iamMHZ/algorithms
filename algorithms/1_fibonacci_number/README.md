## Fibonacci Numbers

Problem is obvious to every one I guess, but if you wanna know more
checkout [this](https://en.wikipedia.org/wiki/Fibonacci_number) wikipedia page.

![problem definition](./docs/0.png)

#### Recursive Solution

+ The **recursive solutions** is very slow and its running time is as demonstrated below:
  ![tn of recursive](./docs/1.png)
+ Why the recursive approach is so slow? Because it recomputes sub-problems over and over again.
  ![why recursive is slow](./docs/2.png)

#### Dynamic Programming Solution

+ Instead of recomputing sub-problems again, store them somewhere and reach them anytime you want them.
+ Basically for computing `F(n)` you need to have last two elements(`F(n-1)` and `F(n-2)`).
  ![optimized approach](./docs/4.png)
+ This algorithm is way too much faster than the recursive approach as its running time is shown.

  ![running time of optimized](./docs/5.png)