# lightweight-blockchain-project-golang

This project is a lightweight blockchain prototype built in Go (**Golang**). It demonstrates the core concepts of a blockchain, including block creation, hashing, and validation. The primary goal of this project is to provide a simple and educational implementation of blockchain fundamentals.

---

## **Features**

- **Genesis Block**: The starting block of the blockchain, initialized with default values.
- **Block Creation**: New blocks can be generated with custom data and linked to the previous block.
- **Hashing**: Blocks are hashed using the **SHA-256** algorithm to ensure immutability.
- **Validation**: The validity of each new block is checked before appending it to the blockchain.

---

## **Getting Started**

### **Prerequisites**

To run this project, ensure you have the following installed on your machine:
- [Go](https://golang.org/) (version 1.19 or later)

---

### **Running the Project**

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/blockchain-prototype.git
   cd blockchain-prototype
   ```

2. Run the program:
   ```bash
   go run main.go
   ```

3. The output will display the blockchain with details for each block, such as index, timestamp, data, previous hash, and current hash.

---

## **Code Overview**

### **Block Structure**

The `Block` struct represents a single block in the blockchain:
```go
type Block struct {
    Index        int
    Timestamp    string
    Data         string
    PreviousHash string
    Hash         string
}
```

### **Functions**

- `calculateHash(block Block)`: Computes the hash for a block using **SHA-256**.
- `generateBlock(previousBlock Block, data string)`: Creates a new block and calculates its hash.
- `isBlockValid(newBlock, previousBlock Block)`: Validates a new block by checking its index, hash, and previous hash.

---

## **Example Output**

When you run the program, you’ll see an output like this:

```plaintext
Index: 0
Timestamp: 2025-01-02 15:04:05 +0000 UTC
Data: Genesis Block
Previous Hash:
Hash: a3b7d1c5469a8e4c8e981b5b1f91b2f3a1b6d23994a6d9b3c5f847a96e9e7a3f

Index: 1
Timestamp: 2025-01-02 15:05:10 +0000 UTC
Data: First Block Data
Previous Hash: a3b7d1c5469a8e4c8e981b5b1f91b2f3a1b6d23994a6d9b3c5f847a96e9e7a3f
Hash: b4e7c1d45c8a6f1b9d2b9481a3f4c2b7e9f82345678a3b9d1f0a6b5c4d2e7c1d

Index: 2
Timestamp: 2025-01-02 15:06:20 +0000 UTC
Data: Second Block Data
Previous Hash: b4e7c1d45c8a6f1b9d2b9481a3f4c2b7e9f82345678a3b9d1f0a6b5c4d2e7c1d
Hash: c9f2b9e8a6f1c3d4e7b2a1f9456789234c5a7b6d8e9f0c3d2a6e7f5b4c1a9e2d
```

---

## **Future Enhancements**

Here are a few ideas for extending this project:
- **Proof-of-Work**: Implement a mining algorithm to simulate computational difficulty.
- **Smart Contracts**: Add functionality for executing simple programmable logic.
- **REST API**: Build a RESTful API to interact with the blockchain.
- **Distributed Network**: Simulate a peer-to-peer network for decentralized functionality.

---

## **License**

This project is open-source and available under the [MIT License](https://opensource.org/licenses/MIT).

---

## **Contact**

Feel free to reach out if you have any questions, suggestions, or feedback!
- **Email**: dineshlokare12@gmail.com
- **LinkedIn**: [LinkedIn Profile](https://www.linkedin.com/in/dinesh-lokare-91855422b/)
  

