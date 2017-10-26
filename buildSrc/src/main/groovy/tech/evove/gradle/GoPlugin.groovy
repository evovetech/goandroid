package tech.evove.gradle

import com.android.build.gradle.api.LibraryVariant
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.api.Task
import org.gradle.api.java.archives.Manifest
import org.gradle.api.tasks.Exec
import org.gradle.api.tasks.bundling.Jar

class GoPlugin implements Plugin<Project> {
    @Override
    void apply(Project project) {
        project.android.libraryVariants.all { LibraryVariant v ->

        }
    }

    static void stuff(Project project, LibraryVariant v) {
        def outputName = project.name
        def srcDir = "${project.projectDir}/src"
        def tmpDir = "${project.buildDir}/tmp/gobind"
        def tmpAar = "${tmpDir}/${outputName}.aar"
        def tmpSrc = "${tmpDir}/${outputName}-sources.jar"
        def cmds = ['gomobile', 'bind', '-v', '-o', "${tmpAar}"]

        Task gobind(type: Exec) {
            inputs.dir(srcDir)
            outputs.file(tmpDir)
            mkdir(tmpDir)

            workingDir srcDir
            commandLine(cmds)
        }

        Task sourcesJar(type: Jar, dependsOn: gobind) {
            group 'artifact'
            description 'creates sources'
            classifier 'sources'
            from zipTree(tmpSrc)
        }

        Task assembleAar(type: Jar, dependsOn: gobind) {
            group 'artifact'
            description 'creates aar'
            extension = "aar"
            Manifest = null
            from zipTree(tmpAar)
        }
    }
}
