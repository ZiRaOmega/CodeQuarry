fetch("api/classement")
  .then((response) => response.json())
  .then((data) => {
    let rank = 0;
    data.forEach((element) => {
      rank++;
      console.log(element);
      const ranking = document.createElement("div");
      ranking.classList.add("ranking");
      ranking.textContent = rank;
      if (rank == 1) {
        ranking.style.color = "gold";
      }
      if (rank == 2) {
        ranking.style.color = "silver";
      }
      if (rank == 3) {
        ranking.style.color = "#cd7f32";
      }

      const classementContainer = document.createElement("div");
      classementContainer.classList.add("classement-container");
      const classementName = document.createElement("p");
      classementName.classList.add("classement-name");
      classementName.textContent = element.LastName;
      const classementFirstName = document.createElement("p");
      classementFirstName.classList.add("classement-firstname");
      classementFirstName.textContent = element.FirstName;
      const classementUsername = document.createElement("p");
      classementUsername.classList.add("classement-username");
      classementUsername.textContent = "@" + element.Username;
      const classementXP = document.createElement("p");
      classementXP.classList.add("classement-XP");
      classementXP.textContent = element.XP.Int64 + "xp";
      const classementRank = document.createElement("p");
      classementRank.classList.add("classement-rank");
      const postLength = document.createElement("p");
      postLength.classList.add("post-length");
      if (element.My_Post != null) {
        postLength.textContent = "Posts: " + element.My_Post.length;
      } else {
        postLength.textContent = "Posts: 0";
      }
      if (element.Rank.String == "Script Kiddie") {
        classementRank.style.color = "Brown";
      }
      if (element.Rank.String == "Bug Hunter") {
        classementRank.style.color = "Silver";
      }
      if (element.Rank.String == "Code Monkey") {
        classementRank.style.color = "Gold";
      }
      if (element.Rank.String == "Git Guru") {
        classementRank.style.color = "Purple";
      }
      if (element.Rank.String == "Stack Overflow Savant") {
        classementRank.style.color = "Red";
      }
      if (element.Rank.String == "Refactoring Rogue") {
        classementRank.style.color = "Yellow";
      }
      if (element.Rank.String == "Agile Archmage") {
        classementRank.style.color = "Green";
      }
      if (element.Rank.String == "Code Whisperer") {
        classementRank.style.color = "Blue";
      }
      if (element.Rank.String == "Heisenbug Debugger") {
        classementRank.style.color = "Orange";
      }
      if (element.Rank.String == "Keyboard Warrior") {
        classementRank.style.color = "Black";
      }
      classementRank.textContent = element.Rank.String;
      classementContainer.appendChild(ranking);
      classementContainer.appendChild(classementUsername);
      classementContainer.appendChild(classementFirstName);
      classementContainer.appendChild(classementName);
      classementContainer.appendChild(classementXP);
      classementContainer.appendChild(classementRank);
      classementContainer.appendChild(postLength);

      document.getElementById("classement").appendChild(classementContainer);
    });
  });
