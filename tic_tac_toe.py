# tic tac toe game for terminal

# make move
# check if square avail
# mark square
# check for win
#	4 directions up, down, left diag, right diag

class Game(object):

	__VALID = set({0,1})

	def __init__(self, current_player, win_num, width, height):
		self.gameboard = Gameboard(width, height)
		self.current_player = current_player
		self.win_num = win_num

	@staticmethod
	def new_game(first_player, width, height, win_num):
		if Game.__is_valid_game(first_player, width, height, win_num):
			return Game(first_player, win_num, width, height)
		print("invalid game configuration... try again cowperson!")

	@staticmethod
	def __is_valid_game(first_player, width, height, win_num):
		if (first_player >= 0 and first_player <= 1) and \
		width > 0 and height > 0 and win_num > 0:
			return True
		return False

	def start(self):
		self.wait_for_input();

	def __wait_for_input(self):
		print("Player: {}".format(self.current_player))
		print("input move in format \"row,col\" e.g. \"3,3\"")
		return (int(x) for x in input().strip().split(","))

	def __switch_player(self):
		if self.current_player == 1:
			self.current_player = 0
		else:
			self.current_player = 1


	def play(self):
		move = self.__wait_for_input()
		row, col = move
		success = self.__make_move(row, col)
		self.gameboard.represent_board()
		if success:
			win = self.gameboard.check_for_win(row, col, self.win_num, self.current_player)
			if win:
				print("player: {} has won the game".format(self.current_player))
			elif not win:
				self.__switch_player()
				self.play()
		if not success:
			print("invalid_move")
			self.play()

	def __make_move(self, row, col):
		if self.gameboard.can_make_move(row, col):
			self.gameboard.mark(row, col, self.current_player)
			return True
		return False

class Gameboard(object):
	def __init__(self, width, height):
		self.board = {}
		self.width = width
		self.height = height

	def check_for_win(self, row, col, win_num, player):
		return self.__check_direction(row, col, 0,1, win_num, player) or self.__check_direction(row, col, 1,1, win_num, player) \
			or self.__check_direction(row, col, 1,0, win_num, player) or self.__check_direction(row, col, 0,0, win_num, player)

	def __check_direction(self, row, col, row_dir, col_dir, win_num, player):
		""" 
		0 , 1 = horizontal
		1, 0 = vertical
		1, 1 = right diag
		0, 0 = left diag
		"""
		if row_dir == 0 and col_dir == 1:
			return self.check_row(row, col, player) >= win_num 
		elif row_dir == 1 and col_dir == 0:
			return self.check_col(row, col, player) >= win_num
		elif row_dir == 1 and col_dir == 1:
			return self.check_right_diag(row, col, player) >= win_num
		return self.check_left_daig(row, col, player) >= win_num

	def check_row(self, row, col, player):
		i = 1
		tally = 1
		while self.player_marked(row + i, col, player):
			tally += 1
			i += 1

		i = 1
		while self.player_marked(row - i, col, player):
			tally += 1
			i -= 1
		return tally

	def check_col(self, row, col, player):
		i = 1
		tally = 1
		while self.player_marked(row, col + i, player):
			tally += 1
			i += 1

		i = 1
		while self.player_marked(row, col - i, player):
			tally -= 1
			i += 1
		return tally

	def check_right_diag(self, row, col, player):
		i = 1
		tally = 1
		while self.player_marked(row - i, col + i, player):
			tally += 1
			i += 1

		i = 1
		while self.player_marked(row + i, col - i, player):
			tally += 1
			i += 1
		return tally

	def check_left_daig(self, row, col, player):
		i = 1
		tally = 1
		while self.player_marked(row + i, col - i, player):
			tally += 1
			i += 1

		i = 1
		while self.player_marked(row - i, col + i, player):
			tally += 1
			i += 1
		return tally

	def player_marked(self, row, col, player):
		# marked squares should be within game board
		# they should also BE in our map AKA unavailable
		if self.__is_valid(row, col) and not self.__is_available(row, col):
			return True if self.__get_value(row, col) == player else False

	def __get_value(self, row, col):
		return self.board[self.__gen_key(row, col)]

	def can_make_move(self, row, col):
		return self.__is_valid(row, col) and self.__is_available(row, col)

	def __is_available(self, row, col):
		return False if self.__gen_key(row, col) in self.board else True

	def __is_valid(self, row, col):
		if (row >= 0 and row < self.height) and (col >= 0 and col < self.width):
			return True
		return False

	@staticmethod
	def __gen_key(row, col):
		return "{},{}".format(row,col)

	def __get_point_tuple(self, point_string):
		return ( int(x) for x in point_string.split(",") )

	def mark(self, row, col, player):
		self.board[self.__gen_key(row, col)] = player

	def represent_board(self):
		""" lets print it
			something like this
			x| |x| |O
			---------
			 | |O| |X 
			---------
			O| | | | 
		"""

		spacer = "".join( ["-"] * ( (self.width * 2) - 1 ) )
		repr_string = ""
		for row in range(self.height):
			# matrix.append([" "] * self.width)
			for col in range(self.width):
				key = self.__gen_key(row, col)
				if key in self.board:
					val = self.board[key]
					if val == 1:
						repr_string += "O"
					else:
						repr_string += "X"
				else:
					repr_string += " "
				if col != self.width - 1:
					repr_string += "|"
			# finished with a row, add newline
			repr_string += "\n"
			# lets also add spacer if were not at the last row
			if row != self.height - 1:
				repr_string += spacer
				repr_string += "\n"

		print(repr_string)


if __name__ == "__main__":
	print("Input \"first_player,width,height,win\" in format \"0, 5,5,3\"")
	print("Player must be 0 or 1")
	first_player, width, height, win = [ int(x) for x in input().strip().split(",") ]

	g = Game.new_game(first_player, width, height, win)
	g.play()


