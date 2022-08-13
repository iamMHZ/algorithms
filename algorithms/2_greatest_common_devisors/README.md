## Greatest Common Divisors

#### Problem Definition

+ Suppose that we have integers `a` and `b`, we need to find the largest integer `d` that divides both `a` and `b`.
  ![Problem Definition](./docs/1.png)

+ Findig GCDs is important in Number Theories and Cryptography.
+ GCD finding algorithms must be efficient enough to run on large numbers.

#### Naive Algorithm:

![Naive Algorithm](./docs/2.png)

+ It is Slow for large numbers and its running time is approximately `a+b`.

#### Euclidean Algorithm

+ Lemma

  ![Lemma](./docs/3.png)
+ Proof

  ![Proof](./docs/4.png)

By computing gdc of `b` and `a'` we do less computation because `a'` is smaller than `a`.

![Algorithms](./docs/5.png)

+ Example

![Example](./docs/6.png)

+ Runtime

![Runtime](./docs/7.png)