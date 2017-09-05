n_o_i = int(raw_input())
f_s = raw_input().split(" ")
s_s = raw_input().split(" ")
t_s = [str(int(x) + int(y)) for x,y in zip(f_s, s_s)]
print " ".join(t_s)
