# Optimized Prime Number Generator in Go

This repository contains an optimized Go implementation for generating prime numbers up to a given limit. The implementation leverages the Sieve of Eratosthenes algorithm and automatically chooses between a non-segmented and segmented approach based on the input size.

## Features
- **Efficient Prime Generation**:
  - For inputs ≤ 10<sup>7</sup>, uses the classic Sieve of Eratosthenes with \( O(n \log \log n) \) complexity.
  - For inputs > 10<sup>7</sup>, employs the segmented sieve to handle large ranges efficiently.

- **Optimizations**:
  - Starts crossing out multiples of each prime from \( i^2 \) to reduce redundant computations.
  - Skips lower non-primes during sieving to save computation time.
  - Implements \( n^2 \) optimization for reducing unnecessary iterations.

- **Output**:
  - Writes all generated prime numbers to a CSV file (`primes_range_<n>.csv`) where `<n>` is the input limit.

## Algorithm: Sieve of Eratosthenes
The Sieve of Eratosthenes is implemented with the following pseudocode:

```text
algorithm Sieve of Eratosthenes is
    input: an integer n > 1.
    output: all prime numbers from 2 through n.

    let A be an array of Boolean values, indexed by integers 2 to n,
    initially all set to true.

    for i = 2, 3, 4, ..., not exceeding √(n) do
        if A[i] is true
            for j = i^2, i^2+i, i^2+2i, i^2+3i, ..., not exceeding n do
                set A[j] := false

    return all i such that A[i] is true.
```
This algorithm achieves \( O(n \log \log n) \) time complexity by efficiently marking non-prime numbers as false.

## Code Structure
- **Non-Segmented Sieve**: Used for smaller ranges (up to 10<sup>7</sup>).
- **Segmented Sieve**: Efficiently handles larger ranges by dividing the range into manageable segments.
- **File Output**: Writes the generated prime numbers to a CSV file with the naming convention `primes_range_<n>.csv`.
- **Time Tracking**: Logs the execution time of the sieve process.

## Download and Use Prime Number Generator

This repository contains precompiled binaries for generating prime numbers. You can find the binaries in the [`bin`](./bin) folder of this repository.

---

### 1. **Locate the Binary**
Choose the binary for your operating system and architecture from the table below:

| Operating System | Architecture | Binary Name                | Path in Repository                |
|-------------------|--------------|----------------------------|------------------------------------|
| Linux            | amd64        | `prime_gen_linux_amd64`   | [`bin/prime_gen_linux_amd64`](./bin/prime_gen_linux_amd64)   |
| macOS            | amd64        | `prime_gen_macos_amd64`   | [`bin/prime_gen_macos_amd64`](./bin/prime_gen_macos_amd64)   |
| Windows          | amd64        | `prime_gen_windows_amd64.exe` | [`bin/prime_gen_windows_amd64.exe`](./bin/prime_gen_windows_amd64.exe) |
| Android          | arm64        | `prime_gen_android_arm64` | [`bin/prime_gen_android_arm64`](./bin/prime_gen_android_arm64) |

---

### 2. **Make the Binary Executable (if needed)**
On Linux, macOS, or Android, you may need to grant execute permissions before running the binary:

```bash
chmod +x <binary_name>
```

Replace `<binary_name>` with the name of the binary you downloaded or copied.

---

### 3. **Run the Binary**
Use the following commands to execute the binary:

#### On Linux/macOS/Android:
```bash
./<binary_name>
```

#### On Windows:
1. Navigate to the `bin` folder and double-click the `.exe` file, or
2. Open a Command Prompt, navigate to the directory containing the `.exe` file, and run:
   ```cmd
   <binary_name>.exe
   ```

---

### 4. **Binary Options**
The binary supports customization via arguments. By default, it generates prime numbers up to a predefined limit. To specify a custom limit, pass the limit as an argument.

#### Example Usage:
```bash
./prime_gen_linux_amd64 1000
```
This command generates prime numbers up to 1000.

#### Available Options:
- `<limit>`: Specify the maximum value for prime number generation.
```

## Usage
1. Build and run the binary.
2. Enter the range when prompted:
   ```bash
   Enter the Range (int): 10000000
   ```
3. The program will choose the appropriate sieve method, generate the primes, and write them to a file named `primes_range_<n>.csv`.
4. It will also display the time taken to execute.

## Performance Statistics
Below is a comparison of execution times for the standard algorithm versus this optimized implementation in Go:

| Limit           | Standard Segmented | This Go Script Segmented | Non-Segmented Standard  | This Go Script Non-Segmented |
|-----------------|-------------------------|------------------------------|---------------------|--------------------------|
| 10,000,000      | 65.346455 ms            | 38.34653 ms                  | 41.40387 ms         | 37.52741 ms             |
| 100,000,000     | 481.153563 ms           | 438.387912 ms                | 884.653602 ms       | 659.169185 ms           |
| 1,000,000,000   | 4.867800215 s           | 3.215119958 s                | 11.220634894 s      | 10.384302883 s          |
| 9,000,000,000   | 29.289063023 s                    | N/A                           | 33.484074665 s      | N/A          |


## References
- [Wikipedia: Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)

