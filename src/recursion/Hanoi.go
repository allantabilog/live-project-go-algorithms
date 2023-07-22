package main

import "fmt"

const numDisks = 3

var posts [][]int

// Add a disk <disk> to the top of the <post>
func push(post []int, disk int) []int {
	return append([]int{disk}, post...)
}

// Remove the top disk from <post>
// Return the disk and the modified post
func pop(post []int) (int, []int) {
	var first = post[0]
	var rest = post[1:]
	return first, rest
}

// move a disk from <fromPost> to <toPost>
func moveDisk(posts [][]int, fromPost, toPost int) {
	var disk int
	disk, posts[fromPost] = pop(posts[fromPost])
	posts[toPost] = push(posts[toPost], disk)
}

// function for drawing the posts
func drawPosts(posts [][]int) {
	// Add 0s to the end of each post so they all have num_disks entries.
	for p := 0; p < 3; p++ {
		for len(posts[p]) < numDisks {
			posts[p] = push(posts[p], 0)
		}
	}

	// Draw the posts.
	for row := 0; row < numDisks; row++ {
		// Draw this row.
		for p := 0; p < 3; p++ {
			// Draw the disk on post p's row.
			fmt.Printf("%d ", posts[p][row])
		}
		fmt.Println()
	}

	// Draw a line between moves.
	fmt.Println("-----")

	// Remove the 0s.
	for p := 0; p < 3; p++ {
		for len(posts[p]) > 0 && posts[p][0] == 0 {
			_, posts[p] = pop(posts[p])
		}
	}
}

func moveDisks(posts [][]int, numToMove, fromPost, toPost, tempPost int) {
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, fromPost, tempPost, toPost)
	}
		moveDisk(posts, fromPost, toPost)
		drawPosts(posts)
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, tempPost, toPost, fromPost)
	}

}

func main() {
	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := numDisks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	drawPosts(posts)

	// Move the disks.
	moveDisks(posts, numDisks, 0, 1, 2)
}
