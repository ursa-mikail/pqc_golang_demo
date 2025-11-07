package main

import (
	"fmt"
	"math/rand"

	"github.com/cloudflare/circl/kem/xwing"
)

func main() {

	seed := make([]byte, 32)
	rand.Read(seed)

	sk, pk := xwing.DeriveKeyPairPacked(seed)

	eseed := make([]byte, 64)
	rand.Read(eseed)

	ss, ct, _ := xwing.Encapsulate(pk, eseed)

	ss2 := xwing.Decapsulate(ct, sk)

	fmt.Printf("Seed for keys %x\n\n", seed)
	fmt.Printf("Seed for encapsulation %x\n\n", eseed)
	fmt.Printf("PK (first 100 bytes) %.200x (size=%d bytes)\n\n", pk, len(pk))
	fmt.Printf("SK %x (size=%d bytes)\n\n", sk, len(sk))

	fmt.Printf("Cipher (first 100 bytes): %.200x (size=%d bytes)\n\n", ct, len(ct))

	fmt.Printf("Shared secret (generated) %x\n\n", ss)
	fmt.Printf("Shared secret (decapsulated) %x\n\n", ss2)

}

/*
% go run main.go
Seed for keys d9d336eb03567e5c7f02a799bc37ac9c7dd5a3952f11e53b14ffb7cea6ec5eb2

Seed for encapsulation 716bfa9dc5510eb0e665e3493d7869ac14c8641b5d88c3725abd09a8776c96ef436aa9da17c69456efd6ec27a94ea2a506b457ad25d62f12045f8fcec129144b

PK (first 100 bytes) 761575890328caeb8646063d8a46228233321fea9fce128b620a0fac2110d8a83761c84faa5375740a1236a45a09e69744426993b35fa60020541165cc88b0bfa2c160431b99917d2bb7ba727ac54c0676636a382c81a55ec00d1e70321edac50e3633df81a7bc41bb8d93c5b158940421a2f6123bd4951feff949f3a19645d1c3ab195b3da98397057829902328235687e443b030919a0173343a21d2667f3e257226cb747b16223711c36b223349f3c77c0c9c8bc073b05a439cd59e430486bd6a2baca00165a0 (size=1216 bytes)

SK d9d336eb03567e5c7f02a799bc37ac9c7dd5a3952f11e53b14ffb7cea6ec5eb2 (size=32 bytes)

Cipher (first 100 bytes): 7197804971627c91aeaed0751201d5cf57f123c58d958cef6ceca97a5aadc708590c396ab95bfb43b9e16b7eda3e82f6450a2e188c9bf8db38b5a9446521a94308ad101a12abcb9d1d899348399368d0fef7a665b6d748aedffd57fc738e370bda27fa587d3584e54d73d3d4495c1c6d7a5329c716a912bc6bc99e1f6422505295e7035fef8685f9d4c36c79008bb9c49a92d14d184fc49216773b7e4395b16a5c23eda39c8916c110a57f616096d07282f976999997aed845d432547e8de171cf93694cda41af40 (size=1120 bytes)

Shared secret (generated) d9e66ac46063e0005daf9bcadbce7b351cf6d86a32c7b99d624dc8bbe90b254e

Shared secret (decapsulated) d9e66ac46063e0005daf9bcadbce7b351cf6d86a32c7b99d624dc8bbe90b254e
*/
