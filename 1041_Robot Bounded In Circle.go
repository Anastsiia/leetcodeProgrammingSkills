/*
On an infinite plane, a robot initially stands at (0, 0) and faces north. Note that:

The north direction is the positive direction of the y-axis.
The south direction is the negative direction of the y-axis.
The east direction is the positive direction of the x-axis.
The west direction is the negative direction of the x-axis.
The robot can receive one of three instructions:

"G": go straight 1 unit.
"L": turn 90 degrees to the left (i.e., anti-clockwise direction).
"R": turn 90 degrees to the right (i.e., clockwise direction).
The robot performs the instructions given in order, and repeats them forever.

Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
*/
package main

type robotPlace struct {
	x    int
	y    int
	view int
}

func (robot *robotPlace) applyMoves(instructions string) {
	for _, v := range instructions {
		switch v {
		case 'G':
			if robot.view == 1 {
				robot.x += 1
			} else if robot.view == 2 {
				robot.y += 1
			} else if robot.view == 3 {
				robot.x -= 1
			} else {
				robot.y -= 1
			}
		case 'L':
			robot.view -= 1
			if robot.view == 0 {
				robot.view = 4
			}
		case 'R':
			robot.view += 1
			if robot.view > 4 {
				robot.view = 1
			}
		}
	}
}

func isRobotBounded(instructions string) bool {
	var robot robotPlace
	robot.x, robot.y = 0, 0
	robot.view = 1
	robot.applyMoves(instructions)
	if (robot.x == 0 && robot.y == 0) || robot.view != 1 {
		return true
	}
	return false
}
