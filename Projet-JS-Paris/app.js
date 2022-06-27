kaboom({
  global: true,
  fullscreen: true,
  scale: 2,
  debug: true,
  clearColor: [0, 0, 1, 0.4],
})

// Speed identifiers
const MOVE_SPEED = 120 // vitesse nous
const JUMP_FORCE = 420 // puissance du saut
const LITTLE_JUMP_FORCE = 280 // quand tu saute sur le mechants petit saut
const BIG_JUMP_FORCE = 450 // saut avec champignon
let CURRENT_JUMP_FORCE = JUMP_FORCE
const FALL_DEATH = 400 // pour savoir quand on meurt de chutes
const ENEMY_SPEED = 20 // vitesse du mechant
let LIVES = 3 // vie de mario


// Game logic
let isJumping = true // si il saute
let isBig = false  // si il est grand
let isOnPipe = false  // si sur un tuyaux
let Direction = false  // si il bouge

//tous les sprites
loadRoot('images/')
loadSprite('link-left','1Xq9biB.png')
loadSprite('cloud', 'Y0oHyRl.png')
loadSprite('coin', 'wbKxhcd.png')
loadSprite('evil-shroom', 'KPO3fR9.png')
loadSprite('brick', 'pogC9x5.png')
loadSprite('block', 'M6rwarW.png')
loadSprite('link-right', 'yZIb8O2.png')
loadSprite('mushroom', '0wMd92p.png')
loadSprite('surprise', 'gesQ1KP.png')
loadSprite('unboxed', 'bdrLpi6.png')
loadSprite('pipe-top-left', 'ReTPiWY.png')
loadSprite('pipe-top-right', 'hj2GK4n.png')
loadSprite('pipe-bottom-left', 'c1cYSbt.png')
loadSprite('pipe-bottom-right', 'nqQ79eI.png')
loadSprite('ground', 'ground.png')
loadSprite('blueGround', 'blueGround.png')
loadSprite('castle', 'supermarcket.png')
loadSprite('metal', 'brownBloc.png')
loadSprite('noBlock', 'noBlock.png')
loadSprite('poto', 'poto.png')
loadSprite('flag', 'flag.png')
loadSprite('back', 'images.png')


loadSprite('top-left', 'ReTPiWY.png')
loadSprite('top-right', 'hj2GK4n.png')
loadSprite('bottom-left', 'c1cYSbt.png')
loadSprite('bottom-right', 'nqQ79eI.png')

loadSprite('blue-block', 'fVscIbn.png')
loadSprite('blue-brick', '3e5YRQd.png')
loadSprite('blue-steel', 'gqVoI2b.png')
loadSprite('blue-evil-shroom', 'SvV4ueD.png')
loadSprite('blue-surprise', 'RMqCc1G.png')

function timer(x){
  for(i=x; i>=0; i--) { (function(temps) {
    setTimeout(function() {console.log(temps);}, (x*1000)-(1000*temps)); })(temps);
  }
}


scene("game", ({ level, score }) => { // dans le jeux
  layers(['bg', 'obj', 'ui','po'], 'obj')
//12éè = fake
  const maps = [ // les maps
    [
      '                                                                                                                                    ',
      '0      ~                     ~                                          ~                     ~                        ~            ',
      '0             ~                                      ~                            ~                                                 ',
      '0                                       ~                      ~                           ¤      ^  ¤    ~                         ',
      '0                     %                                                                     666666666                        7|     ',
      '0                                                                                                                                   ',
      '0                                                6666                     6%6   12                      666*           ==          C',
      '0               %   6*6%6                   12                 ==   =           ()     666                            ===           ',
      '0                                     -+    ()                ===   ==          ()                                   ====           ',
      '0           12                        ()    ()               ====   ===         ==                    12            =====           ',
      '0           ()      ^          ^      ()    ()      ^   ^   =====   ====    ^  ====                   ()        ^  ======     m     ',
      'G                  G                        G                       G                                 G                 G      G',
  
    ],
    [
      '£                                       £',
      '£                                       £',
      '£                                       £',
      '£                                       £',
      '£                                       £',
      '£                                       £',
      '£        @@@@@@              x x        £',
      '£                          x x x        £',//1.2
      'p                        x x x x  x   -+£',
      '£  z    z     z   z z  x x x xzx zx   ()£',
      'g            g      g',
    ], 
    [
      '                                                                                                                                     ',
      '0      ~                     ~                                          ~                     ~                        ~             ',
      '0             ~                                      ~                            ~                                                  ',
      '0                                      ~                      ¤~ ^  z ¤      ¤  ^ ¤             ~                  ~                 ',
      '0                                                              6666666        6666                                                   ',
      '0                                         666       ¤  z   ¤                                                                 7|      ',
      '0                                                    666666            ¤     ¤                                                     C ',//2
      '0         6*6%6          12           ==                                66666        xx         £%£%£                 ==             ',
      '0                        ()          ===        6666                                 xxx                      -+     ===             ',
      '0                        ()         ====                      ¤  z ^ ¤               xxxx                     ()    ====             ',
      '0          z     ^       ()        =====                       666666                xxxxx       z       z    ()   =====      m      ',
      'G                  G                                                                 G                 G                 G           ',
    ],
    [
      '£                                   £',
      '£                  ¤ ^  ^   ¤       £',
      '£                   66666666        £',
      '£           £££                     £',
      '£                                   £',
      '£                                   £',//2.2
      '£        6*66%%          xx         £',
      '£                       xxx         £',
      'p                      xxxx    -+   £',
      '£      z     z   z z  xxxxx    ()   £',
      'g               g          g',
    ], 
    [ //3
    '=                                                          =',
    '=                          %%%                             =',
    '=                                  -+             $        =',
    '=     ¤  ^  ¤                      ()                      =',
    '=      66666   === *     ^   ^     ()    $¤                =',
    '=   *           ~   =G                   =   6666          =',
    '=        ~          =                    =                 =',
    '=                   =                    =                 =',
    '=   =====           =                    =          ====   ',
    '=                   =                    =                                            ',
    '=                   =                    =                                  7|        ',
    '=       ¤   ^ ¤     =                    =                                         C   ',
    '=      ~ 66666      =                    =   ¤  ^  ¤                                  ',
    '=             ~     =                    =    66666                                   ',
    '=                   =                    =                                            ',
    '=                   =                    =                                   m          ',// spwan la
    '=======             =                    =                 G                G',
    '=                   =                    =                 =',
    '=                   =                    =                 =',// dead la
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    '=                   =                    =                 =',
    
  ],
  [
    '£                                 -+£',
    '£                                 ()£',
    '£                  £££££       ££££££',
    '£                                   £',
    '£                                   £',//3.2
    '£       ¤ z z  ¤                    £',
    '£        @@@@@@           xx        £',
    '£                        xxx        £',
    'p                       xxxx        £',
    '£  z    z     z    z   xxxxx        £',
    'g            g      g',
  ],
];

  const levelCfg = { // sprite avec les symboles pour les mettre dans les maps
    width: 20,// largeur
    height: 20,// hauteur
    '~': [sprite('cloud'), scale(0.2)],
    '6': [sprite('brick'),solid(),],
    //'9': [sprite('block')],
    '=': [sprite('block'), solid(), 'bot'],
    '$': [sprite('coin'), 'coin'],
    '%': [sprite('surprise'), solid(), 'coin-surprise'],
    '*': [sprite('surprise'), solid(), 'mushroom-surprise'],
    '}': [sprite('unboxed'), solid()],
    'G': [sprite('ground'), solid(), 'walk'],
    'g': [sprite('blueGround'), solid(), 'walk'],
    'C': [sprite('castle'), scale(0.4)],
    'm': [sprite('metal'), solid(), 'bot'],
    '0': [sprite('noBlock'), solid(), 'bot'],
    '|': [sprite('poto'), solid(), 'fin'],
    '7': [sprite('flag')],
    
    //fake 
    '1': [sprite('top-left'), solid(), scale(0.5), 'walk'],
    '2': [sprite('top-right'), solid(), scale(0.5), 'walk'],
    //true
    '(': [sprite('pipe-bottom-left'), solid(), scale(0.5),'bot'],
    ')': [sprite('pipe-bottom-right'), solid(), scale(0.5),'bot'],
    '3': [sprite('pipe-top-left'), solid(), scale(0.5), 'pipe2'],
    '4': [sprite('pipe-top-right'), solid(), scale(0.5), 'pipe2'],
    '-': [sprite('pipe-top-left'), solid(), scale(0.5), 'pipe'],
    '+': [sprite('pipe-top-right'), solid(), scale(0.5), 'pipe'],
    '^': [sprite('evil-shroom'), solid(), 'dangerous',scale(0.9), body(), {dir: -1},'bot'],
    '#': [sprite('mushroom'), solid(), 'mushroom', body(), {dir: 1},'bot'],
    '!': [sprite('blue-block'), solid(), scale(0.5),'bot'],
    '£': [sprite('blue-brick'), solid(), scale(0.5),'bot'],
    '¤': [sprite('noBlock'), scale(0.5),'bot'],
    'p': [sprite('blue-brick'), solid(), scale(0.5)],
    'z': [sprite('blue-evil-shroom'), solid(), scale(0.40), 'dangerous', body(), {dir: -1},'bot'],
    '@': [sprite('blue-surprise'), solid(), scale(0.5), 'coin-surprise'],
    'x': [sprite('blue-steel'), solid(), scale(0.5),'bot'],

  };

  const gameLevel = addLevel(maps[level], levelCfg); // pour ajouter les sprite

  const player = add([ // pour le perso
    sprite('link-right'), solid(),scale(0.40),
    pos(100, 205),
    body(),
    origin('bot')
  ])

  const scoreLabel = add([ // pour les points
    text(score),
    pos(5, 6),
    layer('ui'),
    {
      value: score,
    }
  ]);
  add([// pour voir les vies que nous avons
    text(LIVES),
    pos(50, 6),
    layer('li')
  ])
  add([// pour voir le temps que nous avons
    text(temps),
    pos(500, 6),
    layer('li')
  ])
  
  
  add([text('level ' + parseInt(level + 1) ), pos(100, 6)]);
  
  function toBig() { // savoir si il est grand et que faire si il est grand
    if (isBig == false) {
      isBig = true;
      player.scale = vec2(0.50);
      CURRENT_JUMP_FORCE = BIG_JUMP_FORCE;
    } else {
      scoreLabel.value += 500;
    };
  } 

  if (isBig) {
    player.scale = vec2(0.50);
    CURRENT_JUMP_FORCE = BIG_JUMP_FORCE;
  }
  
  player.collides('mushroom', (m) => { // pour que le player mange le champignon
    destroy(m); 
    if (isBig == false) {
      toBig();
    }
    score += 500
  });

  function toSmall() { // si il est petit est les capacites de la forme petite
    if (isBig) {
      isBig = false;
      player.scale = vec2(0.4);
      CURRENT_JUMP_FORCE = JUMP_FORCE;
    } else {
      LIVES -= 1
      if (LIVES == 0) {
        go('game_over', { score: scoreLabel.value});
        monAudio.pause();
        // monAudio.play();
        LIVES = 3
      } else {
        go('lose', { score: scoreLabel.value, level: level});
        monAudio.pause();
      }
    };
  }

  player.action(() => { // pour que quand il meurt en tombant il devient petit
    camPos(player.pos);
    if (player.pos.y >= FALL_DEATH) {
      toSmall()
    };
  })
// piece
  player.on("headbump", (obj) => { // pour les bloc avec piece et champignon
    if (obj.is('coin-surprise')) {
      gameLevel.spawn('$', obj.gridPos.sub(0, 1));
      destroy(obj);
      gameLevel.spawn('}', obj.gridPos.sub(0,0));
    };
    if (obj.is('mushroom-surprise')) {
      gameLevel.spawn('#', obj.gridPos.sub(0, 1));
      destroy(obj);
      gameLevel.spawn('}', obj.gridPos.sub(0,0));
    };
  });

  
  action('mushroom', (m) => { // vitesse du champignon
    m.move(m.dir *ENEMY_SPEED, 0)
  });

  collides('mushroom', 'bot', (m) => { // pour que quand il touche un mur il repart en arriere
    m.dir = - m.dir
  })

 

  collides('mushroom', 'dangerous', (m, d) => { // pour que si un mechant et un champignon se tape il partent chaqu'un dans leur direction
    m.dir = - m.dir
    d.dir = - d.dir
  })

  action('dangerous', (d) => { // mouvement du mechant
    d.move(d.dir *ENEMY_SPEED, 0)
  })

  collides('dangerous', 'bot', (d) => {//pour que quand il touche un mur il repart en arriere
    d.dir = - d.dir
  })

  player.collides('dangerous', (d) => {// toute les interractions quand le player touche un mechant
    if (isJumping) {
      scoreLabel.value += 200;
      scoreLabel.text = scoreLabel.value;
      destroy(d);
      player.jump(LITTLE_JUMP_FORCE);
      
    } else {
      toSmall();
      destroy(d);
    };
  });
//piece
  player.collides('coin', (c) => { // interractions quand le player touche une piece
    destroy(c);
    scoreLabel.value++;
    scoreLabel.text = scoreLabel.value;
  });

  player.collides('pipe', () => {// si le player touche un tuyaux pour chanfer de monde
    isOnPipe = true
  })

  player.collides('walk', () => {// annuler le fait qu'il touche un tuyaux
    isOnPipe = false
  })

  player.collides('fin', () => { // touche un drapeau pour la fin
    go('win', { score: scoreLabel.value, level: level});
  })
  
  keyPress('s', () => { // la touche pour changer de monde
    if (isOnPipe) {
      if (level%2 == 0) {
        go('game', {
          level: (level + 1) % maps.length,
          score: scoreLabel.value
        });
      } else {
        go('game', {
          level: (level - 1) % maps.length,
          score: scoreLabel.value
        });
      }
    }
  });
  keyPress('down', () => { // la touche pour changer de monde
    if (isOnPipe) {
      if (level%2 == 0) {
        go('game', {
          level: (level + 1) % maps.length,
          score: scoreLabel.value
        });
      } else {
        go('game', {
          level: (level - 1) % maps.length,
          score: scoreLabel.value
        });
      }
    }
  });

  keyDown('q', () => { // changement de sprite quand il vas a gauche et le fait quil aille a gauche
    if (debug.paused == false) {
      if(isBig == false) {
        player.changeSprite('link-left');
      } else {
        player.changeSprite('link-left');
      }
      player.move(-MOVE_SPEED, 0);
      Direction = true
    }
  });

  keyDown('left', () => { // fleche
    
      if (debug.paused == false) {
        if(isBig == false) {
          player.changeSprite('link-left');
        } else {
          player.changeSprite('link-left');
        }
        player.move(-MOVE_SPEED, 0);
        Direction = true
      }
    
  });
  

  keyDown('d', () => { // = pour la droite
    if (debug.paused == false) {
      if(isBig == false) {
        player.changeSprite('link-right');
      } else {
        player.changeSprite('link-right');
      }
      player.move(MOVE_SPEED, 0);
      Direction = true
    }
  });

  keyDown('right', () => { // fleche
    
      if (debug.paused == false) {
        if(isBig == false) {
          player.changeSprite('link-right');
        } else {
          player.changeSprite('link-right');
        }
        player.move(MOVE_SPEED, 0);
        Direction = true
      }
  
  });

  player.action(() => { // verif si player et en l'air
    if(player.grounded()) {
      isJumping = false;
    }else{
      isJumping = true;
    };
  });


  keyPress('space', () => {// touche pour sauter
    if (player.grounded()) {
      isJumping = true;
      player.jump(CURRENT_JUMP_FORCE);
    };
  });
  keyPress('up', () => {
    if (player.grounded()) {
      isJumping = true;
      player.jump(CURRENT_JUMP_FORCE);
    };
  });
  keyPress('z', () => {
    if (player.grounded()) {
      isJumping = true;
      player.jump(CURRENT_JUMP_FORCE);
    };
  });

  keyPress('p', () => {// mettre en pause
    if (debug.paused == false) {
      debug.paused = true
      monAudio.pause();
      go ("pause", {
        score: scoreLabel.value
        
      });
    } else {
      debug.paused = false
    }
  });

 
});

// toute les scenes avec du texte
scene('win', ({ score, level}) => {
  add([text(score, 32), origin('center'), pos(width()/2, height()/ 2)]);
  add([text("En vrai tranquille t'as un niveau"), origin('center'), pos(width()/2, height(1)/ 5),scale(1.5)]);
  add([text("clique sur espace pour continuer"), origin('center'),pos(width() / 2 , height() / 2 + 180),scale(1.5)]);
  keyPress('space', () => {
    go("game", { level: level+2, score: 0});
  });
});

scene('lose', ({ score, level }) => {
  
  add([text(score, 32), origin('center'), pos(width()/2, height()/ 2)]);
  add([text("\\/"), origin('center'), pos(width()/2, height(1)/ 3+10),scale(1.5)]);
  add([text("||"), origin('center'), pos(width()/2, height(1)/ 3),scale(1.5)]);
  add([text("T'es pas très fort"), origin('center'), pos(width()/2, height(1)/ 5),scale(1.5)]);
  add([text("regarde tes points"), origin('center'), pos(width()/2, height(1)/ 4),scale(1.5)]);
  add([text("recommence si tu pense pouvoir le faire"), origin('center'),pos(width() / 2 , height() / 2 + 180),scale(1.5)]);
  keyPress('space', () => {
    go("game", { level: level, score: 0});
    monAudio.play();
  });
});

scene('game_over', ({ score }) => {
  
  add([text(score, 32), origin('center'), pos(width()/2, height()/ 2)]);
  add([text("\\/"), origin('center'), pos(width()/2, height(1)/ 3+10),scale(1.5)]);
  add([text("||"), origin('center'), pos(width()/2, height(1)/ 3),scale(1.5)]);
  add([text("T'es eclate sa grand mere mon reuf"), origin('center'), pos(width()/2, height(1)/ 5),scale(1.5)]);
  add([text("regarde tes points"), origin('center'), pos(width()/2, height(1)/ 4),scale(1.5)]);
  add([text("recommence si tu pense pouvoir le faire"), origin('center'),pos(width() / 2 , height() / 2 + 180),scale(1.5)]);
  add([text("+ 3 vie vue que tes nul"), origin('center'), pos(width()/2, height()/ 2 + 100),scale(1.5)]);
  keyPress('space', () => {
    go("game", { level: 0, score: 0});
   
  });
});

scene('start', ({}) => {
  add([text("press space to start"), origin('center'), pos(width()/2, height()/ 2)]);
  add([text("Tu as 5 min pour terminer le niveau tous les niveaux"), origin('center'), pos(width()/2, height()/ 3)]);
  add([text("SUPER LINK"), origin('center'), pos(width()/2, height(1)/ 5),scale(2)]);
  keyPress('space', () => {
    go("game", { level: 0, score: 0});
  });
});

scene('pause', ({ score }) => {
  add([text(score, 32), origin('center'), pos(width()/2, height()/ 2)]);
  add([text("Tu veux reprendre ? press r"), origin('center'), pos(width()/2, height(1)/ 5),scale(1.5)]);
  keyPress('r', () => {
    if (debug.paused == true) {
      go("game", { level: 0, score: 0});
      isBig = false
      debug.paused = false
      monAudio.play();
    }
  });

});


start("start", {}, monAudio.play(),timer(400) ); // pour lancer le jeu

/* source: 
https://www.youtube.com/watch?v=2nucjefSr6I
https://www.youtube.com/watch?v=XX93O4ZVUZI
https://kaboomjs.com/
*/