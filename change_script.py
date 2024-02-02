# Read the original file
with open('minitwit.py', 'r') as file:
    content = file.readlines()

# Apply changes
updated_content = []
for line in content:
    if 'from __future__' in line:
        continue  # Skip the __future__ import line
    if 'print ' in line:
        line = line.replace('print ', 'print(')[:-1] + ')\n'  # Update print statement
    updated_content.append(line)

# Write the modified content to a new file
with open('minitwit.py3', 'w') as file:
    file.writelines(updated_content)
