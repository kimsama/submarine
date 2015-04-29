﻿using UnityEngine;
using System;
using Zenject;
using UnityEngine.UI;

namespace Submarine
{
    public class BattleInstaller : MonoInstaller
    {
        [Serializable]
        public class Settings
        {
            [Serializable]
            public class UISettings
            {
                public Button RestartButton;
            }
        }

        public Settings SceneSettings;

        public override void InstallBindings()
        {
            Container.Bind<Settings>().ToSingleInstance(SceneSettings);

            Container.Bind<IInitializable>().ToSingle<TitleController>();
            Container.Bind<BattleController>().ToSingle();
        }
    }
}